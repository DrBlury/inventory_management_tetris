package usecase

import (
	"context"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"linuxcode/inventory_manager/pkg/service/cache"

	"go.uber.org/zap"
)

type appLogicImpl struct {
	queries *repo.Queries
	log     *zap.SugaredLogger
	cache   Cache
}

func NewAppLogic(queries *repo.Queries, logger *zap.SugaredLogger, cache *cache.Cache) appLogicImpl {
	return appLogicImpl{
		queries: queries,
		log:     logger,
		cache:   cache,
	}
}

type InventoryAppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) ([]*domain.InventoryMeta, error)
	AddInventory(ctx context.Context, createInventoryParams *domain.CreateInventoryParams) (*domain.InventoryMeta, error)
	DeleteInventoryById(ctx context.Context, inventoryId int64) error
	GetInventoryById(ctx context.Context, inventoryId int64) (*domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventoryId int64, updateInventoryParams *domain.UpdateInventoryParams) (*domain.InventoryMeta, error)
}

type UserAppLogic interface {
	// Users
	GetUserById(ctx context.Context, userId int64) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	AddUser(ctx context.Context, createUserParams *domain.CreateUserParams) (*domain.User, error)
	DeleteUserById(ctx context.Context, userId int64) error
	GetAllUsers(ctx context.Context) ([]*domain.User, error)
	UpdateUser(ctx context.Context, userId int64, updateUserParams *domain.UpdateUserParams) (*domain.User, error)
}

type ItemAppLogic interface {
	// Items
	GetAllItems(ctx context.Context) ([]*domain.Item, error)
	AddItem(ctx context.Context, createItemParams *domain.CreateItemParams) (*domain.Item, error)
	DeleteItemById(ctx context.Context, itemId int64) error
	GetItemById(ctx context.Context, itemId int64) (*domain.Item, error)
	UpdateItem(ctx context.Context, itemId int64, updateItemParams *domain.UpdateItemParams) error
}

type InventoryItemAppLogic interface {
	// Inventory Item
	AddItemInInventory(ctx context.Context, inventoryId int64, item *domain.Item, quantity int64, durability int64) (*domain.Inventory, error)
	AddItemInInventoryAtPosition(ctx context.Context, inventoryId int64, item *domain.Item, position *domain.Position, quantity int64, durability int64) (*domain.Inventory, error)
	DeleteItemFromInventory(ctx context.Context, inventoryId int64, itemId int64, position *domain.Position, amount int64) (*domain.Inventory, error)
}

type AppLogic interface {
	InventoryAppLogic
	UserAppLogic
	ItemAppLogic
	InventoryItemAppLogic
}
