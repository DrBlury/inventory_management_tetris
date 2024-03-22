package usecase

import (
	"context"
	"fmt"
	"linuxcode/inventory_manager/pkg/domain"

	"go.uber.org/zap"
)

// AddItemInInventory adds an item to the inventory at the first possible position
func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int, item domain.Item, quantity int, durability int) error {
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
func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int, item domain.Item, position domain.Position, quantity int, durability int) error {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return err
	}

	// add the item to the inventory
	inventoryItem, err := inventory.AddItemAtPosition(item, &position, quantity, durability)
	if err != nil {
		return err
	}
	a.log.Info("added item to inventory", zap.Int("inventoryId", inventoryId), zap.Int("itemId", inventoryItem.Item.ItemMeta.ID), zap.Any("position", position))

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
func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int, itemId int, position domain.Position, amount int) error {
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
