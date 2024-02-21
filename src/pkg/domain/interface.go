package domain

import (
	"context"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
)

type appLogicImpl struct {
	queries repo.Queries
}

func NewAppLogic(queries repo.Queries) appLogicImpl {
	return appLogicImpl{
		queries: queries,
	}
}

type AppLogic interface {
	// Inventories
	GetAllInventories(ctx context.Context) ([]Inventory, error)
	AddInventory(ctx context.Context, inventory Inventory) error
	DeleteInventoryById(ctx context.Context, inventoryId int) error
	GetInventoryById(ctx context.Context, inventoryId int) (Inventory, error)
	AddItemInInventory(ctx context.Context, inventoryId int, item Item) error
	AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item Item, position Position) error
	DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position Position, amount int) error

	// Items
	GetAllItems(ctx context.Context) ([]Item, error)
	AddItem(ctx context.Context, item Item) error
	DeleteItemById(ctx context.Context, itemId int) error
	GetItemById(ctx context.Context, itemId int) (Item, error)

	// Users
	GetUserById(ctx context.Context, userId int) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	AddUser(ctx context.Context, user User) error
	DeleteUserById(ctx context.Context, userId int) error
}

// Inventories
func (a appLogicImpl) GetAllInventories(ctx context.Context) ([]Inventory, error) {
	_, err := a.queries.ListInventories(ctx)
	if err != nil {
		return nil, err
	}

	// transform to domain model
	// TODO
	return nil, nil
}

func (a appLogicImpl) AddInventory(ctx context.Context, inventory Inventory) error {
	panic("not implemented") // TODO: Implement
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

func (a appLogicImpl) AddItem(ctx context.Context, item Item) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) DeleteItemById(ctx context.Context, itemId int) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) GetItemById(ctx context.Context, itemId int) (Item, error) {
	panic("not implemented") // TODO: Implement
}

// Users
func (a appLogicImpl) GetUserById(ctx context.Context, userId int) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) GetUserByUsername(ctx context.Context, username string) (User, error) {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) AddUser(ctx context.Context, user User) error {
	panic("not implemented") // TODO: Implement
}

func (a appLogicImpl) DeleteUserById(ctx context.Context, userId int) error {
	panic("not implemented") // TODO: Implement
}
