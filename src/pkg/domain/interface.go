package domain

import (
	"context"
	"encoding/json"
	"fmt"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"linuxcode/inventory_manager/pkg/service/cache"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type appLogicImpl struct {
	queries *repo.Queries
	log     *zap.SugaredLogger
	cache   *cache.Cache
}

func NewAppLogic(queries *repo.Queries, logger *zap.SugaredLogger, cache *cache.Cache) appLogicImpl {
	return appLogicImpl{
		queries: queries,
		log:     logger,
	}
}

type AppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) ([]*InventoryMeta, error)
	AddInventory(ctx context.Context, createInventoryParams CreateInventoryParams) (*Inventory, error)
	DeleteInventoryById(ctx context.Context, inventoryId int) error
	GetInventoryById(ctx context.Context, inventoryId int) (*Inventory, error)
	AddItemInInventory(ctx context.Context, inventoryId int, item Item) error
	AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error
	DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error
	UpdateInventory(ctx context.Context, inventoryId int, updateInventoryParams CreateInventoryParams) error

	// Items
	GetAllItems(ctx context.Context) ([]*Item, error)
	AddItem(ctx context.Context, createItemParams CreateItemParams) error
	DeleteItemById(ctx context.Context, itemId int) error
	GetItemById(ctx context.Context, itemId int) (*Item, error)
	UpdateItem(ctx context.Context, itemId int, updateItemParams CreateItemParams) error

	// Users
	GetUserById(ctx context.Context, userId int) (*User, error)
	GetUserByUsername(ctx context.Context, username string) (*User, error)
	AddUser(ctx context.Context, createUserParams CreateUserParams) (*User, error)
	DeleteUserById(ctx context.Context, userId int) error
	GetAllUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, userId int, updateUserParams CreateUserParams) error

	// TODO all update functions are missing
}

func (a appLogicImpl) setItemInCache(ctx context.Context, itemID int, item *Item) error {
	// marshal the item into a json
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}
	itemIDString := strconv.Itoa(itemID)
	key := fmt.Sprint("itemID-", itemIDString)
	return a.cache.SetString(ctx, key, string(itemJSON))
}

func (a appLogicImpl) getItemFromCache(ctx context.Context, itemID int) (*Item, error) {
	itemIDString := strconv.Itoa(itemID)
	key := fmt.Sprint("itemID-", itemIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		return nil, err
	}
	// Unmarshal the value into a domain.Item
	var item Item
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (a appLogicImpl) setInventoryInCache(ctx context.Context, inventoryID int, inventory *Inventory) error {
	// marshal the inventory into a json
	inventoryJSON, err := json.Marshal(inventory)
	if err != nil {
		return err
	}
	itemIDString := strconv.Itoa(inventoryID)
	key := fmt.Sprint("inventoryID-", itemIDString)
	return a.cache.SetString(ctx, key, string(inventoryJSON))
}

func (a appLogicImpl) getInventoryFromCache(ctx context.Context, inventoryID int) (*Inventory, error) {
	inventoryIDString := strconv.Itoa(inventoryID)
	key := fmt.Sprint("inventoryID-", inventoryIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		return nil, err
	}
	// Unmarshal the value into a domain.Item
	var inventory Inventory
	err = json.Unmarshal([]byte(val), &inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}

// GetAllInventories returns metadata of all inventories
// TODO add pagination
func (a appLogicImpl) GetAllInventories(ctx context.Context) ([]*InventoryMeta, error) {
	// Try to hit cache
	allInventories, err := a.cache.GetString(ctx, "allInventories")
	if err == nil {
		// We got a cache hit! Wonderful!
		var domainInventories []*InventoryMeta
		err = json.Unmarshal([]byte(allInventories), &domainInventories)
		if err == nil {
			return domainInventories, err
		} else {
			a.log.Error("error unmarshalling all inventories from json", zap.Error(err))
		}
	}
	repoInventories, err := a.queries.ListInventories(ctx)
	if err != nil {
		return nil, err
	}
	// map repo model to domain model
	var inventoryMetas = make([]*InventoryMeta, 0, len(repoInventories))
	for i, rInv := range repoInventories {
		inventoryMetas[i] = MapRepoInventoryToDomainInventoryMeta(&rInv)
	}

	// store all inventories in cache

	// marshal to json
	inventoryMetasJSON, err := json.Marshal(inventoryMetas)
	if err != nil {
		return inventoryMetas, err
	}
	err = a.cache.SetString(ctx, "allInventories", string(inventoryMetasJSON))
	if err != nil {
		return inventoryMetas, err
	}

	return inventoryMetas, nil
}

// AddInventory creates a new inventory without any items for the given user
func (a appLogicImpl) AddInventory(ctx context.Context, createInventoryParams CreateInventoryParams) (*InventoryMeta, error) {
	// log user the inventory is created for
	a.log.Info("creating inventory for user", zap.Int("userId", createInventoryParams.UserID))
	createdInventory, err := a.queries.CreateInventory(ctx, repo.CreateInventoryParams{
		UserID:    pgtype.Int4{Int32: int32(createInventoryParams.UserID), Valid: true},
		Invname:   pgtype.Text{String: createInventoryParams.Name, Valid: true},
		Width:     pgtype.Int4{Int32: int32(createInventoryParams.Width), Valid: true},
		Height:    pgtype.Int4{Int32: int32(createInventoryParams.Height), Valid: true},
		MaxWeight: pgtype.Int4{Int32: int32(createInventoryParams.MaxWeight), Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// Map repo model to domain model
	invMeta := MapRepoInventoryToDomainInventoryMeta(&createdInventory)

	// log what inventory was created
	a.log.Info("created inventory", zap.String("name", invMeta.Name))

	return invMeta, nil
}

// DeleteInventoryById deletes the inventory with the given id
func (a appLogicImpl) DeleteInventoryById(ctx context.Context, inventoryId int) error {
	deletedInventory, err := a.queries.DeleteInventory(ctx, int32(inventoryId))
	if err != nil {
		return err
	}

	// log what inventory was deleted
	a.log.Info("deleted inventory", zap.String("name", deletedInventory.Invname.String))
	return nil
}

// GetInventoryById returns the inventory with the given id
func (a appLogicImpl) GetInventoryById(ctx context.Context, inventoryId int) (*Inventory, error) {
	// Get inventory from cache
	domainInventory, err := a.getInventoryFromCache(ctx, inventoryId)
	if err == nil {
		// We got a cache hit! Wonderful!
		return domainInventory, nil
	}

	repoInventory, err := a.queries.GetInventory(ctx, int32(inventoryId))
	if err != nil {
		return nil, err
	}

	// get all items in the inventory
	repoInventoryItems, err := a.queries.ListInventoryItems(ctx, pgtype.Int4{Int32: int32(inventoryId), Valid: true})
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainInventoryItems := MapRepoInventoryItemsToDomainInventoryItems(&repoInventoryItems)

	domainInventory = &Inventory{
		InventoryMeta: InventoryMeta{
			ID:        int(repoInventory.ID),
			Name:      repoInventory.Invname.String,
			Width:     int(repoInventory.Width.Int32),
			Height:    int(repoInventory.Height.Int32),
			MaxWeight: int(repoInventory.MaxWeight.Int32),
		},
		Items: *domainInventoryItems,
	}

	// populate the InventoryItems with data about the items
	for i, invItem := range domainInventory.Items {
		item, err := a.GetItemById(ctx, invItem.Item.ItemMeta.ID)
		if err != nil {
			return nil, err
		}
		// Map the missing fields from the item to the inventory item
		domainInventory.Items[i].Item.ItemMeta = ItemMeta{
			ID:       item.ItemMeta.ID,
			Shape:    item.ItemMeta.Shape,
			Weight:   item.ItemMeta.Weight,
			MaxStack: item.ItemMeta.MaxStack,
		}
	}

	// store the inventory in the cache
	err = a.setInventoryInCache(ctx, inventoryId, domainInventory)
	if err != nil {
		a.log.Error("error setting inventory in cache", zap.Error(err))
	}

	return domainInventory, nil
}

// AddItemInInventory adds an item to the inventory at the first possible position
func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int, item Item, quantity int, durability int) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	inventory.AddItem(item, quantity, durability)

	// TODO update the inventory in the database and invalidate cache for this inventory
	return nil
}

// AddItemInInventoryAtPosition adds an item to the inventory at the given position
func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	// add the item to the inventory
	inventoryItem, err := inventory.AddItemAtPosition(item, &position, 1, 100)
	if err != nil {
		return err
	}
	a.log.Info("added item to inventory", zap.Int("inventoryId", inventoryId), zap.Int("itemId", inventoryItem.Item.ItemMeta.ID), zap.Any("position", position))

	// TODO update the inventory in the database and invalidate cache for this inventory
	return nil
}

// DeleteItemFromInventory deletes the given amount of items from the inventory at the given position
func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	// find the item in the inventory
	for i, invItem := range inventory.Items {
		if invItem.Item.ItemMeta.ID == itemId && invItem.Position == position {
			// check the amount of that item in the inventory
			if invItem.Quantity < amount {
				return fmt.Errorf("not enough items in inventory")
			}

			// decrease the amount of that item in the inventory
			inventory.Items[i].Quantity -= amount

			// if the amount is 0, remove the item from the inventory
			if inventory.Items[i].Quantity == 0 {
				inventory.Items = append(inventory.Items[:i], inventory.Items[i+1:]...)
			}
			break
		}
	}

	// TODO update the inventory in the database and invalidate cache for this inventory
	return nil
}

// GetAllItems returns all items that exist
func (a appLogicImpl) GetAllItems(ctx context.Context) (*[]Item, error) {
	// check for cache hit
	allItems, err := a.cache.GetString(ctx, "allItems")
	if err == nil {
		// We got a cache hit! Wonderful!
		var domainItems []Item
		err = json.Unmarshal([]byte(allItems), &domainItems)
		if err != nil {
			a.log.Error("error unmarshalling all items from json", zap.Error(err))
			return &domainItems, err
		}
		return &domainItems, nil
	}

	repoItems, err := a.queries.ListItems(ctx)
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainItems := MapRepoItemsToDomainItems(repoItems...)

	// turn ALL to json and store in cache as one
	jsonAllItems, err := json.Marshal(domainItems)
	if err != nil {
		a.log.Error("error marshalling all items to json", zap.Error(err))
		return domainItems, err
	}
	err = a.cache.SetString(ctx, "allItems", string(jsonAllItems))
	if err != nil {
		a.log.Error("error setting all items in cache", zap.Error(err))
		return domainItems, err
	}

	return domainItems, nil
}

// AddItem adds a new item to the database
func (a appLogicImpl) AddItem(ctx context.Context, createItemParams CreateItemParams) error {
	repoItemType := repo.ItemType(createItemParams.Type)
	createdItem, err := a.queries.CreateItem(ctx, repo.CreateItemParams{
		Name:        pgtype.Text{String: createItemParams.Name, Valid: true},
		Description: pgtype.Text{String: createItemParams.Description, Valid: true},
		Variant:     pgtype.Text{String: createItemParams.Variant, Valid: true},
		Type:        repo.NullItemType{ItemType: repoItemType, Valid: true},
		BuyValue:    pgtype.Int4{Int32: int32(createItemParams.BuyValue), Valid: true},
		SellValue:   pgtype.Int4{Int32: int32(createItemParams.SellValue), Valid: true},
		MaxStack:    pgtype.Int4{Int32: int32(createItemParams.MaxStack), Valid: true},
		Weight:      pgtype.Int4{Int32: int32(createItemParams.Weight), Valid: true},
		Durability:  pgtype.Int4{Int32: int32(createItemParams.Durability), Valid: true},
		Height:      pgtype.Int4{Int32: int32(createItemParams.Shape.Height), Valid: true},
		Width:       pgtype.Int4{Int32: int32(createItemParams.Shape.Width), Valid: true},
		Rawshape:    pgtype.Text{String: createItemParams.Shape.RawShape, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	// log what item was created
	a.log.Info("created item", zap.String("name", createdItem.Name.String))
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.Error("error invalidating allItems in cache", zap.Error(err))
	}

	err = a.cache.Invalidate(ctx, strconv.Itoa(int(createdItem.ID)))
	if err != nil {
		a.log.Error("error invalidating item in cache", zap.Error(err))
	}

	return nil
}

// DeleteItemById deletes the item with the given id
func (a appLogicImpl) DeleteItemById(ctx context.Context, itemId int) error {
	repoItem, err := a.queries.DeleteItem(ctx, int32(itemId))
	if err != nil {
		return err
	}

	// log what item was deleted
	a.log.Info("deleted item", zap.String("name", repoItem.Name.String))

	// invalidate cache for this item
	err = a.cache.Invalidate(ctx, strconv.Itoa(itemId))
	if err != nil {
		a.log.Error("error invalidating item in cache", zap.Error(err))
	}

	// invalidate cache for all items
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.Error("error invalidating allItems in cache", zap.Error(err))
	}

	return nil
}

// GetItemById returns the item with the given id
func (a appLogicImpl) GetItemById(ctx context.Context, itemId int) (*Item, error) {
	result, err := a.getItemFromCache(ctx, itemId)
	if err == nil {
		// We got a cache hit! Wonderful!
		return result, nil
	}

	repoItem, err := a.queries.GetItem(ctx, int32(itemId))
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainItems := *MapRepoItemsToDomainItems(repoItem)
	domainItem := &domainItems[0]

	// Store the item in the cache, ignore error for now
	err = a.setItemInCache(ctx, itemId, domainItem)
	if err != nil {
		a.log.Error("error setting item in cache", zap.Error(err))
	}

	return domainItem, nil
}

// GetAllUsers returns all users that exist
func (a appLogicImpl) GetAllUsers(ctx context.Context) ([]User, error) {
	users, err := a.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	var domainUsers = make([]User, 0, len(users))
	for i, rUser := range users {
		user := User{
			ID:       int(rUser.ID),
			Username: rUser.Username.String,
			Email:    rUser.Email.String,
		}
		domainUsers[i] = user
	}
	return domainUsers, nil
}

// GetUserById returns the user with the given id
func (a appLogicImpl) GetUserById(ctx context.Context, userId int) (User, error) {
	user, err := a.queries.GetUser(ctx, int32(userId))
	if err != nil {
		return User{}, err
	}

	// map repo model to domain model
	domainUser := User{
		ID:       int(user.ID),
		Username: user.Username.String,
		Email:    user.Email.String,
	}
	return domainUser, nil
}

// GetUserByUsername returns the user with the given username
func (a appLogicImpl) GetUserByUsername(ctx context.Context, username string) (User, error) {
	user, err := a.queries.GetUserByUsername(ctx, pgtype.Text{String: username, Valid: true})
	if err != nil {
		return User{}, err
	}

	// map repo model to domain model
	domainUser := User{
		ID:       int(user.ID),
		Username: user.Username.String,
		Email:    user.Email.String,
	}
	return domainUser, nil
}

// HashPassword hashes the given password using bcrypt and the given salt
func HashPassword(password, salt string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

// AddUser creates a new user
func (a appLogicImpl) AddUser(ctx context.Context, createUserParams CreateUserParams) (*repo.User, error) {
	// hash password using bcrypt
	salt := "saltySalt-" + createUserParams.Username
	hashedPassword, err := HashPassword(createUserParams.Password, salt)
	if err != nil {
		return nil, err
	}

	createdUser, err := a.queries.CreateUser(ctx, repo.CreateUserParams{
		Username:     pgtype.Text{String: createUserParams.Username, Valid: true},
		Email:        pgtype.Text{String: createUserParams.Email, Valid: true},
		Salt:         pgtype.Text{String: salt, Valid: true},
		PasswordHash: pgtype.Text{String: hashedPassword, Valid: true},
		CreatedAt:    pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// log what user was created
	a.log.Info("created user", zap.String("username", createUserParams.Username))
	return &createdUser, nil
}

// DeleteUserById deletes the user with the given id
func (a appLogicImpl) DeleteUserById(ctx context.Context, userId int) error {
	repoUser, err := a.queries.DeleteUser(ctx, int32(userId))
	if err != nil {
		return err
	}

	// log what user was deleted
	a.log.Info("deleted user", zap.String("username", repoUser.Username.String))
	return nil
}
