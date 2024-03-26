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
	cache   *cache.Cache
}

func NewAppLogic(queries *repo.Queries, logger *zap.SugaredLogger, cache *cache.Cache) appLogicImpl {
	return appLogicImpl{
		queries: queries,
		log:     logger,
		cache:   cache,
	}
}

type AppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) (*[]domain.InventoryMeta, error)
	AddInventory(ctx context.Context, createInventoryParams domain.CreateInventoryParams) (*domain.InventoryMeta, error)
	DeleteInventoryById(ctx context.Context, inventoryId int) error
	GetInventoryById(ctx context.Context, inventoryId int) (*domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventoryId int, updateInventoryParams domain.UpdateInventoryParams) (*domain.InventoryMeta, error)

	// Items
	GetAllItems(ctx context.Context) (*[]domain.Item, error)
	AddItem(ctx context.Context, createItemParams domain.CreateItemParams) error
	DeleteItemById(ctx context.Context, itemId int) error
	GetItemById(ctx context.Context, itemId int) (*domain.Item, error)
	UpdateItem(ctx context.Context, itemId int, updateItemParams domain.UpdateItemParams) error

	// Users
	GetUserById(ctx context.Context, userId int) (*domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
	AddUser(ctx context.Context, createUserParams domain.CreateUserParams) (*domain.User, error)
	DeleteUserById(ctx context.Context, userId int) error
	GetAllUsers(ctx context.Context) (*[]domain.User, error)
	UpdateUser(ctx context.Context, userId int, updateUserParams domain.UpdateUserParams) (*domain.User, error)

	// Inventory Item
	AddItemInInventory(ctx context.Context, inventoryId int, item domain.Item, quantity int, durability int) error
	AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item domain.Item, position domain.Position, quantity int, durability int) error
	DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position domain.Position, amount int) error
}
