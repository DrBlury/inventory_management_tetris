package usecase

import (
	"context"
	"encoding/json"
	"linuxcode/inventory_manager/pkg/domain"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"go.uber.org/zap"
)

// GetAllItems returns all items that exist
func (a appLogicImpl) GetAllItems(ctx context.Context) (*[]domain.Item, error) {
	// check for cache hit
	allItems, err := a.cache.GetString(ctx, "allItems")
	if err == nil {
		// We got a cache hit! Wonderful!
		var domainItems []domain.Item
		err = json.Unmarshal([]byte(allItems), &domainItems)
		if err != nil {
			a.log.Error("error unmarshalling all items from json", zap.Error(err))
			return &domainItems, err
		}
		return &domainItems, nil
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
		a.log.Error("error marshalling all items to json", zap.Error(err))
		return domainItems, err
	}
	err = a.cache.SetString(ctx, "allItems", string(jsonAllItems))
	if err != nil {
		a.log.Error("error setting all items in cache", zap.Error(err))
		return domainItems, err
	}

	return domainItems, nil
}

// AddItem adds a new item to the database
func (a appLogicImpl) AddItem(ctx context.Context, createItemParams domain.CreateItemParams) error {
	repoItemType := repo.ItemType(createItemParams.Type)
	createdItem, err := a.queries.CreateItem(ctx, repo.CreateItemParams{
		Name:        pgtype.Text{String: createItemParams.Name, Valid: true},
		Description: pgtype.Text{String: createItemParams.Description, Valid: true},
		Variant:     pgtype.Text{String: createItemParams.Variant, Valid: true},
		Type:        repo.NullItemType{ItemType: repoItemType, Valid: true},
		BuyValue:    pgtype.Int4{Int32: int32(createItemParams.BuyValue), Valid: true},
		SellValue:   pgtype.Int4{Int32: int32(createItemParams.SellValue), Valid: true},
		MaxStack:    pgtype.Int4{Int32: int32(createItemParams.MaxStack), Valid: true},
		Weight:      pgtype.Int4{Int32: int32(createItemParams.Weight), Valid: true},
		Durability:  pgtype.Int4{Int32: int32(createItemParams.Durability), Valid: true},
		Height:      pgtype.Int4{Int32: int32(createItemParams.Shape.Height), Valid: true},
		Width:       pgtype.Int4{Int32: int32(createItemParams.Shape.Width), Valid: true},
		Rawshape:    pgtype.Text{String: createItemParams.Shape.RawShape, Valid: true},
		CreatedAt:   pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		return err
	}

	// log what item was created
	a.log.Info("created item", zap.String("name", createdItem.Name.String))
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.Error("error invalidating allItems in cache", zap.Error(err))
	}

	err = a.cache.Invalidate(ctx, strconv.Itoa(int(createdItem.ID)))
	if err != nil {
		a.log.Error("error invalidating item in cache", zap.Error(err))
	}

	return nil
}

// DeleteItemById deletes the item with the given id
func (a appLogicImpl) DeleteItemById(ctx context.Context, itemId int) error {
	repoItem, err := a.queries.DeleteItem(ctx, int32(itemId))
	if err != nil {
		return err
	}

	// log what item was deleted
	a.log.Info("deleted item", zap.String("name", repoItem.Name.String))

	// invalidate cache for this item
	err = a.cache.Invalidate(ctx, strconv.Itoa(itemId))
	if err != nil {
		a.log.Error("error invalidating item in cache", zap.Error(err))
	}

	// invalidate cache for all items
	err = a.cache.Invalidate(ctx, "allItems")
	if err != nil {
		a.log.Error("error invalidating allItems in cache", zap.Error(err))
	}

	return nil
}

// GetItemById returns the item with the given id
func (a appLogicImpl) GetItemById(ctx context.Context, itemId int) (*domain.Item, error) {
	result, err := a.getItemFromCache(ctx, itemId)
	if err == nil {
		// We got a cache hit! Wonderful!
		return result, nil
	}

	repoItem, err := a.queries.GetItem(ctx, int32(itemId))
	if err != nil {
		return nil, err
	}

	// map repo model to domain model
	domainItems := *domain.MapRepoItemsToDomainItems(repoItem)
	domainItem := &domainItems[0]

	// Store the item in the cache, ignore error for now
	err = a.setItemInCache(ctx, itemId, domainItem)
	if err != nil {
		a.log.Error("error setting item in cache", zap.Error(err))
	}

	return domainItem, nil
}
