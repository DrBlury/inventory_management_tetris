package dboTransform

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
)

type DBOTransform interface {
	// To domain model
	ToItem(*repo.Item) *domain.Item
	ToInventoryItem(*repo.InventoryItem) *domain.InventoryItem
	ToInventory(*repo.Inventory) *domain.Inventory
	ToUser(*repo.User) *domain.User

	// To repo model
	ToRepoItem(*domain.Item) *repo.Item
	ToRepoInventoryItem(*domain.InventoryItem) *repo.InventoryItem
	ToRepoInventory(*domain.Inventory) *repo.Inventory
	ToRepoUser(*domain.User) *repo.User
}

type dboTransformImpl struct{}

func New() DBOTransform {
	return dboTransformImpl{}
}

// === MAP REPO TO DOMAIN MODEL ===
func (dboTransformImpl) ToItem(item *repo.Item) *domain.Item {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToInventoryItem(inventoryItem *repo.InventoryItem) *domain.InventoryItem {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToInventory(inventory *repo.Inventory) *domain.Inventory {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToUser(user *repo.User) *domain.User {
	// TODO implement
	return nil
}

// === MAP DOMAIN TO REPO MODEL ===
func (dboTransformImpl) ToRepoItem(item *domain.Item) *repo.Item {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToRepoInventoryItem(inventoryItem *domain.InventoryItem) *repo.InventoryItem {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToRepoInventory(inventory *domain.Inventory) *repo.Inventory {
	// TODO implement
	return nil
}

func (dboTransformImpl) ToRepoUser(user *domain.User) *repo.User {
	// TODO implement
	return nil
}
