package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"linuxcode/inventory_manager/pkg/domain"
	"strconv"
)

func (a appLogicImpl) setItemInCache(ctx context.Context, itemID int, item *domain.Item) error {
	// marshal the item into a json
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}
	itemIDString := strconv.Itoa(itemID)
	key := fmt.Sprint("itemID-", itemIDString)
	return a.cache.SetString(ctx, key, string(itemJSON))
}

func (a appLogicImpl) getItemFromCache(ctx context.Context, itemID int) (*domain.Item, error) {
	itemIDString := strconv.Itoa(itemID)
	key := fmt.Sprint("itemID-", itemIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		return nil, err
	}
	// Unmarshal the value into a domain.Item
	var item domain.Item
	err = json.Unmarshal([]byte(val), &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (a appLogicImpl) setInventoryInCache(ctx context.Context, inventoryID int, inventory *domain.Inventory) error {
	// marshal the inventory into a json
	inventoryJSON, err := json.Marshal(inventory)
	if err != nil {
		return err
	}
	itemIDString := strconv.Itoa(inventoryID)
	key := fmt.Sprint("inventoryID-", itemIDString)
	return a.cache.SetString(ctx, key, string(inventoryJSON))
}

func (a appLogicImpl) getInventoryFromCache(ctx context.Context, inventoryID int) (*domain.Inventory, error) {
	inventoryIDString := strconv.Itoa(inventoryID)
	key := fmt.Sprint("inventoryID-", inventoryIDString)
	val, err := a.cache.GetString(ctx, key)
	if err != nil {
		return nil, err
	}
	// Unmarshal the value into a domain.Item
	var inventory domain.Inventory
	err = json.Unmarshal([]byte(val), &inventory)
	if err != nil {
		return nil, err
	}

	return &inventory, nil
}
