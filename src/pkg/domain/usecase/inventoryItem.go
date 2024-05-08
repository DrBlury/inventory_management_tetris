package usecase

import (
	"context"
	"fmt"
	domain "linuxcode/inventory_manager/pkg/domain/model"

	"go.uber.org/zap"
)

// AddItemInInventory adds an item to the inventory at the first possible position
func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int64, item *domain.Item, quantity int64, durability int64) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	inventory.AddItem(item, quantity, durability)

	// TODO update the inventory in the database

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}
	return nil
}

// AddItemInInventoryAtPosition adds an item to the inventory at the given position
func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int64, item *domain.Item, position *domain.Position, quantity int64, durability int64) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	// add the item to the inventory
	inventoryItem, err := inventory.AddItemAtPosition(item, position, quantity, durability)
	if err != nil {
		return err
	}
	a.log.Info("added item to inventory", zap.Int64("inventoryId", inventoryId), zap.Int64("itemId", inventoryItem.Item.ItemMeta.Id), zap.Any("position", position))

	// TODO update the inventory in the database

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}

	return nil
}

// DeleteItemFromInventory deletes the given amount of items from the inventory at the given position
func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int64, itemId int64, position *domain.Position, amount int64) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	// find the item in the inventory
	for i, invItem := range inventory.Items {
		currentPos := invItem.Position
		if invItem.Item.ItemMeta.Id == itemId && currentPos == position {
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
