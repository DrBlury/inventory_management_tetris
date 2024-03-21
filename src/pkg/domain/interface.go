package domain

import (
	"context"
	"encoding/json"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"linuxcode/inventory_manager/pkg/service/cache"
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

func NewAppLogic(queries *repo.Queries, logger *zap.SugaredLogger) appLogicImpl {
	return appLogicImpl{
		queries: queries,
		log:     logger,
	}
}

type AppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) ([]InventoryMeta, error)
	AddInventory(ctx context.Context, createInventoryParams CreateInventoryParams) (*repo.Inventory, error)
	DeleteInventoryById(ctx context.Context, inventoryId int) error
	GetInventoryById(ctx context.Context, inventoryId int) (Inventory, error)
	AddItemInInventory(ctx context.Context, inventoryId int, item Item) error
	AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error
	DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error

	// Items
	GetAllItems(ctx context.Context) ([]Item, error)
	AddItem(ctx context.Context, createItemParams CreateItemParams) error
	DeleteItemById(ctx context.Context, itemId int) error
	GetItemById(ctx context.Context, itemId int) (Item, error)

	// Users
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	AddUser(ctx context.Context, createUserParams CreateUserParams) (*repo.User, error)
	DeleteUserById(ctx context.Context, userId int) error
	GetAllUsers(ctx context.Context) ([]User, error)
}

func (a appLogicImpl) SetItemInCache(ctx context.Context, itemID string, item Item) error {
	// marshal the item into a json
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}
	return a.cache.SetString(ctx, itemID, string(itemJSON))
}

func (a appLogicImpl) GetItemFromCache(ctx context.Context, itemID string) (Item, error) {
	val, err := a.cache.GetString(ctx, itemID)
	if err != nil {
		return Item{}, err
	}
	// Unmarshal the value into a domain.Item
	var item Item
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

// GetAllInventories returns metadata of all inventories
// TODO add pagination
func (a appLogicImpl) GetAllInventories(ctx context.Context) ([]InventoryMeta, error) {
	repoInventories, err := a.queries.ListInventories(ctx)
	if err != nil {
		return nil, err
	}
	// map repo model to domain model
	var inventoryMetas = make([]InventoryMeta, 0, len(repoInventories))
	for i, rInv := range repoInventories {
		InventoryMeta := InventoryMeta{
			ID:        int(rInv.ID),
			Name:      rInv.Invname.String,
			Width:     int(rInv.Width.Int32),
			Height:    int(rInv.Height.Int32),
			MaxWeight: int(rInv.MaxWeight.Int32),
		}
		inventoryMetas[i] = InventoryMeta
	}

	return inventoryMetas, nil
}

// AddInventory creates a new inventory without any items for the given user
func (a appLogicImpl) AddInventory(ctx context.Context, createInventoryParams CreateInventoryParams) (*repo.Inventory, error) {
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
	return &createdInventory, nil
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
	var domainInventoryItems = make([]InventoryItem, 0, len(repoInventoryItems))
	for i, rItem := range repoInventoryItems {
		itemId := int(rItem.ItemID.Int32)
		item, err := a.GetItemById(ctx, itemId)
		if err != nil {
			continue // skip this item, it doesn't exist?
		}

		domainInventoryItems[i] = InventoryItem{
			Item: item,
			Position: Position{
				X:        int(rItem.PositionX.Int32),
				Y:        int(rItem.PositionY.Int32),
				Rotation: int(rItem.Rotation.Int32),
			},
			Quantity:       int(rItem.Quantity.Int32),
			DurabilityLeft: int(rItem.DurabilityLeft.Int32),
		}
	}
	domainInventory := Inventory{
		InventoryMeta: InventoryMeta{
			ID:        int(repoInventory.ID),
			Name:      repoInventory.Invname.String,
			Width:     int(repoInventory.Width.Int32),
			Height:    int(repoInventory.Height.Int32),
			MaxWeight: int(repoInventory.MaxWeight.Int32),
		},
		Items: domainInventoryItems,
	}

	return &domainInventory, nil
}

// AddItemInInventory adds an item to the inventory at the first possible position
func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int, item Item) error {
	panic("not implemented") // TODO: Implement
}

// AddItemInInventoryAtPosition adds an item to the inventory at the given position
func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error {
	panic("not implemented") // TODO: Implement
}

// DeleteItemFromInventory deletes the given amount of items from the inventory at the given position
func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error {
	panic("not implemented") // TODO: Implement
}

// GetAllItems returns all items that exist
func (a appLogicImpl) GetAllItems(ctx context.Context) ([]Item, error) {
	panic("not implemented") // TODO: Implement
}

// AddItem adds a new item to the database
func (a appLogicImpl) AddItem(ctx context.Context, createItemParams CreateItemParams) error {
	// TODO maybe item needs a weight variable???
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
	return nil
}

// GetItemById returns the item with the given id
func (a appLogicImpl) GetItemById(ctx context.Context, itemId int) (Item, error) {
	result, err := a.GetItemFromCache(ctx, string(itemId))
	if err == nil {
		// We got a cache hit! Wonderful!
		return result, nil
	}

	repoItem, err := a.queries.GetItem(ctx, int32(itemId))
	if err != nil {
		return Item{}, err
	}

	// map repo model to domain model

	// TODO check if they are named the same or if we need explicit mapping
	domainItemType := ItemType(repoItem.Type.ItemType)
	domainItem := Item{
		ID:          int(repoItem.ID),
		Name:        repoItem.Name.String,
		Description: repoItem.Description.String,
		Type:        domainItemType,
		Variant:     repoItem.Variant.String,
		BuyValue:    int(repoItem.BuyValue.Int32),
		SellValue:   int(repoItem.SellValue.Int32),
		MaxStack:    int(repoItem.MaxStack.Int32),
		Weight:      int(repoItem.Weight.Int32),
		Durability:  int(repoItem.Durability.Int32),
		Shape: Shape{
			Height:   int(repoItem.Height.Int32),
			Width:    int(repoItem.Width.Int32),
			RawShape: repoItem.Rawshape.String,
		},
	}

	// Store the item in the cache, ignore error for now
	err = a.SetItemInCache(ctx, string(itemId), domainItem)
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
