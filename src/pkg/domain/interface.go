package domain

import (
	"context"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type appLogicImpl struct {
	queries *repo.Queries
	log     *zap.SugaredLogger
}

func NewAppLogic(queries *repo.Queries, logger *zap.SugaredLogger) appLogicImpl {
	return appLogicImpl{
		queries: queries,
		log:     logger,
	}
}

type AppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) ([]repo.Inventory, error)
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

// Inventories
func (a appLogicImpl) GetAllInventories(ctx context.Context) ([]repo.Inventory, error) {
	inventories, err := a.queries.ListInventories(ctx)
	if err != nil {
		return nil, err
	}

	// transform to domain model
	// TODO
	return inventories, nil
}

func (a appLogicImpl) AddInventory(ctx context.Context, createInventoryParams CreateInventoryParams) (*repo.Inventory, error) {
	// log user the inventory is created for
	a.log.Info("creating inventory for user", zap.Int("userId", createInventoryParams.UserID))
	createdInventory, err := a.queries.CreateInventory(ctx, repo.CreateInventoryParams{
		UserID:    pgtype.Int4{Int32: int32(createInventoryParams.UserID), Valid: true},
		Invname:   pgtype.Text{String: createInventoryParams.Name, Valid: true},
		SizeH:     pgtype.Int4{Int32: int32(createInventoryParams.Width), Valid: true},
		SizeV:     pgtype.Int4{Int32: int32(createInventoryParams.Height), Valid: true},
		MaxWeight: pgtype.Int4{Int32: int32(createInventoryParams.MaxWeight), Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}
	return &createdInventory, nil
}

func (a appLogicImpl) DeleteInventoryById(ctx context.Context, inventoryId int) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) GetInventoryById(ctx context.Context, inventoryId int) (Inventory, error) {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int, item Item) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error {
	panic("not implemented") // TODO: Implement
}

// Items
func (a appLogicImpl) GetAllItems(ctx context.Context) ([]Item, error) {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) AddItem(ctx context.Context, createItemParams CreateItemParams) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) DeleteItemById(ctx context.Context, itemId int) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) GetItemById(ctx context.Context, itemId int) (Item, error) {
	panic("not implemented") // TODO: Implement
}

// Users
func (a appLogicImpl) GetAllUsers(ctx context.Context) ([]User, error) {
	panic("not implemented") // TODO: Implement
}
func (a appLogicImpl) GetUserById(ctx context.Context, userId int) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) GetUserByUsername(ctx context.Context, username string) (User, error) {
	panic("not implemented") // TODO: Implement
}

func HashPassword(password, salt string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

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
	return &createdUser, nil
}

func (a appLogicImpl) DeleteUserById(ctx context.Context, userId int) error {
	panic("not implemented") // TODO: Implement
}
