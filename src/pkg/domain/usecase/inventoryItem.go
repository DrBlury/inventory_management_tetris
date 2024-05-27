package usecase

import (
	"context"
	"fmt"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joomcode/errorx"
	"go.uber.org/zap"
)

// AddItemInInventory adds an item to the inventory at the first possible position
func (a appLogicImpl) AddItemInInventory(ctx context.Context, inventoryId int64, item *domain.Item, quantity int64, durability int64) (*domain.Inventory, error) {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not exist.", inventoryId)
	}
	positionSuggestion, ok := inventory.AddItem(item, quantity, durability)
	if ok {
		// TODO update the inventory in the database
		a.queries.AddItemToInventory(ctx, repo.AddItemToInventoryParams{
			InventoryID: pgtype.Int4{Int32: int32(inventory.InventoryMeta.Id), Valid: true},
			ItemID:      pgtype.Int4{Int32: int32(item.ItemMeta.Id), Valid: true},
			PositionX:   pgtype.Int4{Int32: int32(positionSuggestion.X), Valid: true},
			PositionY:   pgtype.Int4{Int32: int32(positionSuggestion.Y), Valid: true},
			Rotation:    pgtype.Int4{Int32: int32(positionSuggestion.Rotation), Valid: true},
			Quantity:    pgtype.Int4{Int32: int32(quantity), Valid: true},
			// TODO missing!?
			//Durability: durability
		})
	}

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}
	return inventory, nil
}

// AddItemInInventoryAtPosition adds an item to the inventory at the given position
func (a appLogicImpl) AddItemInInventoryAtPosition(ctx context.Context, inventoryId int64, item *domain.Item, position *domain.Position, quantity int64, durability int64) (*domain.Inventory, error) {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not exist.", inventoryId)
	}

	// add the item to the inventory
	inventoryItem, err := inventory.AddItemAtPosition(item, position, quantity, durability)
	if err != nil {
		return nil, errorx.IllegalArgument.New("Addint the item with id %d at PosX: %d, PosY: %d, Rot: %d was not possible: %s", item.ItemMeta.Id, position.X, position.Y, position.Rotation, err.Error())
	}
	a.log.Info("added item to inventory", zap.Int64("inventoryId", inventoryId), zap.Int64("itemId", inventoryItem.Item.ItemMeta.Id), zap.Any("position", position))

	positionSuggestion, ok := inventory.AddItem(item, quantity, durability)
	if ok {
		// TODO update the inventory in the database
		a.queries.AddItemToInventory(ctx, repo.AddItemToInventoryParams{
			InventoryID: pgtype.Int4{Int32: int32(inventory.InventoryMeta.Id), Valid: true},
			ItemID:      pgtype.Int4{Int32: int32(item.ItemMeta.Id), Valid: true},
			PositionX:   pgtype.Int4{Int32: int32(positionSuggestion.X), Valid: true},
			PositionY:   pgtype.Int4{Int32: int32(positionSuggestion.Y), Valid: true},
			Rotation:    pgtype.Int4{Int32: int32(positionSuggestion.Rotation), Valid: true},
			Quantity:    pgtype.Int4{Int32: int32(quantity), Valid: true},
			// TODO missing!?
			//Durability: durability
		})
	}

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}

	return inventory, nil
}

// DeleteItemFromInventory deletes the given amount of items from the inventory at the given position
func (a appLogicImpl) DeleteItemFromInventory(ctx context.Context, inventoryId int64, itemId int64, position *domain.Position, amount int64) (*domain.Inventory, error) {
	// get inventory
	inventory, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		// TODO move this error inside GetInventoryById and decorate here! Also do so in the other calling methods...
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not exist.", inventoryId)
	}

	foundItem := false

	// find the item in the inventory
	for i, invItem := range inventory.Items {
		currentPos := invItem.Position
		if invItem.Item.ItemMeta.Id == itemId && currentPos == position {
			foundItem = true
			// check the amount of that item in the inventory
			if invItem.Quantity < amount {
				return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not have the amount to be removed: %d", inventory.InventoryMeta.Id, amount)
			}

			// decrease the amount of that item in the inventory
			inventory.Items[i].Quantity -= amount
			// TODO decrease the amount of the item in the database

			// if the amount is 0, remove the item from the inventory
			if inventory.Items[i].Quantity == 0 {
				inventory.Items = append(inventory.Items[:i], inventory.Items[i+1:]...)
				// TODO Remove the item completely from the database
			}
			break
		}
	}

	if !foundItem {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not have item to be removed at PosX: %d, PosY: %d, Rot: %d", inventory.InventoryMeta.Id, position.X, position.Y, position.Rotation)
	}

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}
	return inventory, nil
}

// MoveItemWithinInventory moves an item within an inventory from startPos to Endpos.
// The amount to move is optional, if not set it will move the whole stack of items.
func (a appLogicImpl) MoveItemWithinInventory(ctx context.Context, inventoryId int64, itemId int64, startPos *domain.Position, endpos *domain.Position, amount int64) (*domain.Inventory, error) {
	// get inventory
	inv, err := a.GetInventoryById(ctx, inventoryId)
	if err != nil {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not exist.", inventoryId)
	}
	a.log.Info("Found inventory with ID: ", zap.Int64("inventoryId", inv.InventoryMeta.Id))

	// check if the item is in the inventory
	itemExists := false
	var itemToMove *domain.InventoryItem
	for _, itemInInv := range inv.Items {
		if itemInInv.Item.ItemMeta.Id == itemId && itemInInv.Position == startPos {
			itemExists = true
			itemToMove = itemInInv
			break
		}
	}

	if !itemExists {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not have item to be moved at PosX: %d, PosY: %d, Rot: %d", inventoryId, startPos.X, startPos.Y, startPos.Rotation)
	}

	a.log.Info("Found item to move with ID: ", zap.Int64("itemId", itemToMove.Item.ItemMeta.Id))

	// check if the amount to move is valid
	if amount < itemToMove.Quantity {
		return nil, errorx.IllegalArgument.New("The requested Inventory with Id: %d does not have the amount to be moved: %d", inventoryId, amount)
	}

	// move the item
	a.queries.UpdateInventoryItem(ctx, repo.UpdateInventoryItemParams{
		ID:          int32(itemToMove.Item.ItemMeta.Id),
		InventoryID: pgtype.Int4{Int32: int32(inv.InventoryMeta.Id), Valid: true},
		ItemID:      pgtype.Int4{Int32: int32(itemToMove.Item.ItemMeta.Id), Valid: true},
		PositionX:   pgtype.Int4{Int32: int32(endpos.X), Valid: true},
		PositionY:   pgtype.Int4{Int32: int32(endpos.Y), Valid: true},
		Rotation:    pgtype.Int4{Int32: int32(endpos.Rotation), Valid: true},
		Quantity:    pgtype.Int4{Int32: int32(amount), Valid: true},
	})

	// invalidate cache
	key := fmt.Sprint("inventoryID-", inventoryId)
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}
	return nil, nil
}
