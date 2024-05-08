package usecase

import (
	"context"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

// GetAllInventories returns metadata of all inventories
// TODO add pagination
func (a appLogicImpl) GetAllInventories(ctx context.Context) ([]*domain.InventoryMeta, error) {
	allInventories, err := a.getInventoriesFromCache(context.Background())
	if err == nil && allInventories != nil {
		return allInventories, nil
	}

	a.log.Info("getting all inventories from database instead of cache")

	repoInventories, err := a.queries.ListInventories(ctx)
	if err != nil {
		return nil, err
	}
	// map repo model to domain model
	var inventoryMetas = make([]*domain.InventoryMeta, 0, len(repoInventories))
	for _, rInv := range repoInventories {
		inventoryMetas = append(inventoryMetas, domain.MapRepoInventoryToDomainInventoryMeta(&rInv))
	}

	// store all inventories in cache
	err = a.setInventoriesInCache(ctx, inventoryMetas)
	if err != nil {
		a.log.Error("error setting all inventories in cache", zap.Error(err))
	}

	return inventoryMetas, nil
}

// AddInventory creates a new inventory without any items for the given user
func (a appLogicImpl) AddInventory(ctx context.Context, createInventoryParams *domain.CreateInventoryParams) (*domain.InventoryMeta, error) {
	// log user the inventory is created for
	a.log.Info("creating inventory for user", zap.Int64("userId", createInventoryParams.UserId))
	createdInventory, err := a.queries.CreateInventory(ctx, repo.CreateInventoryParams{
		UserID:    pgtype.Int4{Int32: int32(createInventoryParams.UserId), Valid: true},
		Invname:   pgtype.Text{String: createInventoryParams.Name, Valid: true},
		Width:     pgtype.Int4{Int32: int32(createInventoryParams.Width), Valid: true},
		Height:    pgtype.Int4{Int32: int32(createInventoryParams.Height), Valid: true},
		MaxWeight: pgtype.Int4{Int32: int32(createInventoryParams.MaxWeight), Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// invalidate cache
	err = a.cache.Invalidate(ctx, "allInventoriesMeta")
	if err != nil {
		a.log.Error("error invalidating all inventories in cache", zap.Error(err))
	}

	// Map repo model to domain model
	invMeta := domain.MapRepoInventoryToDomainInventoryMeta(&createdInventory)

	// log what inventory was created
	a.log.Info("created inventory", zap.String("name", invMeta.Name))

	return invMeta, nil
}

// DeleteInventoryById deletes the inventory with the given id
func (a appLogicImpl) DeleteInventoryById(ctx context.Context, inventoryId int64) error {
	deletedInventory, err := a.queries.DeleteInventory(ctx, int32(inventoryId))
	if err != nil {
		return err
	}

	// invalidate cache
	err = a.cache.Invalidate(ctx, "allInventoriesMeta")
	if err != nil {
		a.log.Error("error invalidating all inventories in cache", zap.Error(err))
	}

	// log what inventory was deleted
	a.log.Info("deleted inventory", zap.String("name", deletedInventory.Invname.String))
	return nil
}

// GetInventoryById returns the inventory with the given id
func (a appLogicImpl) GetInventoryById(ctx context.Context, inventoryId int64) (*domain.Inventory, error) {
	// Get inventory from cache
	domainInventory, err := a.getInventoryFromCache(ctx, inventoryId)
	if err == nil && domainInventory != nil {
		// We got a cache hit! Wonderful!
		return domainInventory, nil
	}

	repoInventory, err := a.queries.GetInventory(ctx, int32(inventoryId))
	if err != nil {
		return nil, err
	}

	// get all items in the inventory
	repoInventoryItems, err := a.queries.ListInventoryItems(ctx, pgtype.Int4{Int32: int32(inventoryId), Valid: true})
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainInventoryItems := domain.MapRepoInventoryItemsToDomainInventoryItems(&repoInventoryItems)

	domainInventory = &domain.Inventory{
		InventoryMeta: &domain.InventoryMeta{
			Id:        int64(repoInventory.ID),
			Name:      repoInventory.Invname.String,
			Width:     int64(repoInventory.Width.Int32),
			Height:    int64(repoInventory.Height.Int32),
			MaxWeight: int64(repoInventory.MaxWeight.Int32),
		},
		Items: domainInventoryItems,
	}

	// populate the InventoryItems with data about the items
	for i, invItem := range domainInventory.Items {
		item, err := a.GetItemById(ctx, invItem.Item.ItemMeta.Id)
		if err != nil {
			return nil, err
		}
		// Map the missing fields from the item to the inventory item
		domainInventory.Items[i].Item.ItemMeta = &domain.ItemMeta{
			Id:       item.ItemMeta.Id,
			Shape:    item.ItemMeta.Shape,
			Weight:   item.ItemMeta.Weight,
			MaxStack: item.ItemMeta.MaxStack,
		}
	}

	// store the inventory in the cache
	err = a.setInventoryInCache(ctx, inventoryId, domainInventory)
	if err != nil {
		a.log.Error("error setting inventory in cache", zap.Error(err))
	}

	return domainInventory, nil
}

// UpdateInventory updates the inventory with the given id
func (a appLogicImpl) UpdateInventory(ctx context.Context, inventoryId int64, updateInventoryParams *domain.UpdateInventoryParams) (*domain.InventoryMeta, error) {
	updatedInventory, err := a.queries.UpdateInventory(ctx, repo.UpdateInventoryParams{
		UserID:    pgtype.Int4{Int32: int32(updateInventoryParams.UserId), Valid: true},
		Invname:   pgtype.Text{String: updateInventoryParams.Name, Valid: true},
		Width:     pgtype.Int4{Int32: int32(updateInventoryParams.Width), Valid: true},
		Height:    pgtype.Int4{Int32: int32(updateInventoryParams.Height), Valid: true},
		MaxWeight: pgtype.Int4{Int32: int32(updateInventoryParams.MaxWeight), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// invalidate cache
	key := "inventoryID-" + strconv.Itoa(int(inventoryId))
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.Error("error invalidating inventory in cache", zap.String("key", key), zap.Error(err))
	}

	err = a.cache.Invalidate(ctx, "allInventories")
	if err != nil {
		a.log.Error("error invalidating all inventories in cache", zap.String("key", key), zap.Error(err))
	}

	// map repo model to domain model
	invMeta := domain.MapRepoInventoryToDomainInventoryMeta(&updatedInventory)

	return invMeta, nil
}
