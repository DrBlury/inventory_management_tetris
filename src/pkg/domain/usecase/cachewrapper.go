package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	"strconv"

	"go.uber.org/zap"
)

// get items from cache
func (a appLogicImpl) getAllItems(ctx context.Context) ([]*domain.Item, error) {
	// check for cache hit
	allItems, err := a.cache.GetString(ctx, "allItems")
	if err != nil {
		a.log.Error("error getting all items from cache", zap.Error(err))
	}
	if allItems == "" {
		return nil, nil
	}
	// unmarshal the items
	var domainItems []*domain.Item
	err = json.Unmarshal([]byte(allItems), &domainItems)
	if err != nil {
		a.log.Error("error unmarshalling all items from json", zap.Error(err))
		return domainItems, err
	}
	return domainItems, nil
}
func (a appLogicImpl) setItemInCache(ctx context.Context, itemID int64, item *domain.Item) error {
	// marshal the item into a json
	itemJSON, err := json.Marshal(item)
	if err != nil {
		a.log.Error("error marshalling item", zap.Error(err))
		return err
	}
	itemIDString := strconv.Itoa(int(itemID))
	key := fmt.Sprint("itemID-", itemIDString)
	err = a.cache.SetString(ctx, key, string(itemJSON))
	if err != nil {
		a.log.Error("error setting item in cache", zap.Error(err))
		return err
	}
	return nil
}

func (a appLogicImpl) getItemFromCache(ctx context.Context, itemID int64) (*domain.Item, error) {
	itemIDString := strconv.Itoa(int(itemID))
	key := fmt.Sprint("itemID-", itemIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		a.log.Error("error getting item from cache", zap.Error(err))
	}
	if val == "" {
		return nil, nil
	}
	// Unmarshal the value into a domain.Item
	var item domain.Item
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		a.log.Error("error unmarshalling item from cache", zap.Error(err))
		return nil, err
	}

	return &item, nil
}

func (a appLogicImpl) setInventoryInCache(ctx context.Context, inventoryID int64, inventory *domain.Inventory) error {
	// marshal the inventory into a json
	inventoryJSON, err := json.Marshal(inventory)
	if err != nil {
		a.log.Error("error marshalling inventory", zap.Error(err))
		return err
	}
	itemIDString := strconv.Itoa(int(inventoryID))
	key := fmt.Sprint("inventoryID-", itemIDString)
	err = a.cache.SetString(ctx, key, string(inventoryJSON))
	if err != nil {
		a.log.Error("error setting inventory in cache", zap.Error(err))
		return err
	}
	return nil
}

func (a appLogicImpl) getInventoryFromCache(ctx context.Context, inventoryID int64) (*domain.Inventory, error) {
	inventoryIDString := strconv.Itoa(int(inventoryID))
	key := fmt.Sprint("inventoryID-", inventoryIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		a.log.Error("error getting inventory from cache", zap.Error(err))
		return nil, err
	}
	if val == "" {
		return nil, nil
	}
	// Unmarshal the value into a domain.Item
	var inventory domain.Inventory
	err = json.Unmarshal([]byte(val), &inventory)
	if err != nil {
		a.log.Error("error unmarshalling inventory from cache", zap.Error(err))
		return nil, err
	}

	return &inventory, nil
}

func (a appLogicImpl) setInventoriesInCache(ctx context.Context, inventories []*domain.InventoryMeta) error {
	// marshal the inventory into a json
	inventoriesJSON, err := json.Marshal(inventories)
	if err != nil {
		a.log.Error("error marshalling inventories", zap.Error(err))
		return err
	}
	key := "allInventoriesMeta"
	err = a.cache.SetString(ctx, key, string(inventoriesJSON))
	if err != nil {
		a.log.Error("error setting inventories in cache", zap.Error(err))
		return err
	}
	return nil
}

func (a appLogicImpl) getInventoriesFromCache(ctx context.Context) ([]*domain.InventoryMeta, error) {
	key := "allInventoriesMeta"
	val, err := a.cache.GetString(ctx, key)
	// switch on the error to handle it
	if err != nil {
		a.log.Error("error getting inventories from cache", zap.Error(err))
		return nil, err
	}
	if val == "" {
		return nil, nil
	}
	// Unmarshal the value into a domain.Item
	var inventories []*domain.InventoryMeta
	err = json.Unmarshal([]byte(val), &inventories)
	if err != nil {
		a.log.Error("error unmarshalling inventories from cache", zap.Error(err))
		return nil, err
	}

	return inventories, nil
}
