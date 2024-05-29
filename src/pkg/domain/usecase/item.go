package usecase

import (
	"context"
	"encoding/json"
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	dboTransform "linuxcode/inventory_manager/pkg/repo/transform"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

// GetAllItems returns all items that exist
func (a appLogicImpl) GetAllItems(ctx context.Context) ([]*domain.Item, error) {
	// check for cache hits
	allItems, err := a.getAllItems(ctx)
	if err == nil && allItems != nil {
		return allItems, nil
	}

	repoItems, err := a.queries.ListItems(ctx)
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainItems := domain.MapRepoItemsToDomainItems(repoItems...)

	// turn ALL to json and store in cache as one
	jsonAllItems, err := json.Marshal(domainItems)
	if err != nil {
		a.log.With(zap.Error(err)).Error("error marshalling all items to json")
		return domainItems, err
	}
	err = a.cache.SetString(ctx, "allItems", string(jsonAllItems))
	if err != nil {
		a.log.With(zap.Error(err)).Error("error setting all items in cache")
		return domainItems, err
	}

	return domainItems, nil
}

// AddItem adds a new item to the database
func (a appLogicImpl) AddItem(ctx context.Context, createItemParams *domain.CreateItemParams) (*domain.Item, error) {
	a.log.With(zap.String("name", createItemParams.Name)).Info("creating item")
	a.log.With(zap.Any("type", createItemParams.Type)).Info("item type")

	repoItemType := dboTransform.ToRepoItemType(createItemParams.Type)
	a.log.With(zap.String("repo type", string(repoItemType))).Info("item type")
	createdItem, err := a.queries.CreateItem(ctx, repo.CreateItemParams{
		Name:       pgtype.Text{String: createItemParams.Name, Valid: true},
		Text:       pgtype.Text{String: createItemParams.Text, Valid: true},
		Variant:    pgtype.Text{String: createItemParams.Variant, Valid: true},
		Type:       repo.NullItemType{ItemType: repoItemType, Valid: true},
		BuyValue:   pgtype.Int4{Int32: int32(createItemParams.BuyValue), Valid: true},
		SellValue:  pgtype.Int4{Int32: int32(createItemParams.SellValue), Valid: true},
		MaxStack:   pgtype.Int4{Int32: int32(createItemParams.MaxStack), Valid: true},
		Weight:     pgtype.Int4{Int32: int32(createItemParams.Weight), Valid: true},
		Durability: pgtype.Int4{Int32: int32(createItemParams.Durability), Valid: true},
		Height:     pgtype.Int4{Int32: int32(createItemParams.Shape.Height), Valid: true},
		Width:      pgtype.Int4{Int32: int32(createItemParams.Shape.Width), Valid: true},
		Rawshape:   pgtype.Text{String: createItemParams.Shape.RawShape, Valid: true},
		CreatedAt:  pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return nil, err
	}

	// log what item was created
	a.log.With(zap.Int32("name", createdItem.ID)).Info("created item")
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating allItems in cache")
	}

	key := "ItemID-" + strconv.Itoa(int(createdItem.ID))
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating item in cache")
	}

	// map repo model to domain model
	var domainItem *domain.Item
	domainItems := domain.MapRepoItemsToDomainItems(createdItem)
	if len(domainItems) > 0 {
		domainItem = domainItems[0]
	}

	return domainItem, nil
}

// DeleteItemById deletes the item with the given id
func (a appLogicImpl) DeleteItemById(ctx context.Context, itemId int64) error {
	repoItem, err := a.queries.DeleteItem(ctx, int32(itemId))
	if err != nil {
		return err
	}

	// log what item was deleted
	a.log.With(zap.Int32("id", repoItem.ID)).Info("deleted item")

	// invalidate cache for this item
	key := "ItemID-" + strconv.Itoa(int(repoItem.ID))
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating item in cache")
	}

	// invalidate cache for all items
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating allItems in cache")
	}

	return nil
}

// GetItemById returns the item with the given id
func (a appLogicImpl) GetItemById(ctx context.Context, itemId int64) (*domain.Item, error) {
	result, err := a.getItemFromCache(ctx, itemId)
	if err == nil && result != nil {
		// We got a cache hit! Wonderful!
		return result, nil
	}

	repoItem, err := a.queries.GetItem(ctx, int32(itemId))
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainItems := domain.MapRepoItemsToDomainItems(repoItem)
	domainItem := domainItems[0]

	// Store the item in the cache, ignore error for now
	err = a.setItemInCache(ctx, itemId, domainItem)
	if err != nil {
		a.log.With(zap.Error(err)).Error("error setting item in cache")
	}

	return domainItem, nil
}

func (a appLogicImpl) UpdateItem(ctx context.Context, itemId int64, updateItemParams *domain.UpdateItemParams) error {
	repoItemType := repo.ItemType(updateItemParams.Type)
	_, err := a.queries.UpdateItem(ctx, repo.UpdateItemParams{
		ID:         int32(itemId),
		Name:       pgtype.Text{String: updateItemParams.Name, Valid: true},
		Text:       pgtype.Text{String: updateItemParams.Text, Valid: true},
		Variant:    pgtype.Text{String: updateItemParams.Variant, Valid: true},
		Type:       repo.NullItemType{ItemType: repoItemType, Valid: true},
		BuyValue:   pgtype.Int4{Int32: int32(updateItemParams.BuyValue), Valid: true},
		SellValue:  pgtype.Int4{Int32: int32(updateItemParams.SellValue), Valid: true},
		MaxStack:   pgtype.Int4{Int32: int32(updateItemParams.MaxStack), Valid: true},
		Weight:     pgtype.Int4{Int32: int32(updateItemParams.Weight), Valid: true},
		Durability: pgtype.Int4{Int32: int32(updateItemParams.Durability), Valid: true},
		Height:     pgtype.Int4{Int32: int32(updateItemParams.Shape.Height), Valid: true},
		Width:      pgtype.Int4{Int32: int32(updateItemParams.Shape.Width), Valid: true},
		Rawshape:   pgtype.Text{String: updateItemParams.Shape.RawShape, Valid: true},
	})
	if err != nil {
		return err
	}

	// log what item was updated
	a.log.With(zap.Int64("id", itemId)).Info("updated item")
	key := "ItemID-" + strconv.Itoa(int(itemId))
	err = a.cache.Invalidate(ctx, key)
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating item in cache")
	}

	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.With(zap.Error(err)).Error("error invalidating allItems in cache")
	}

	return nil
}
