package dtoTransform

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

// === MAP DTO ENUM TO DOMAIN ENUM ===
func ToItemType(itemType server.ItemType) domain.ItemType {
	switch itemType {
	case server.MeleeWeapon:
		return domain.ItemType_MELEE_WEAPON
	case server.RangedWeapon:
		return domain.ItemType_RANGED_WEAPON
	case server.ConsumableWeapon:
		return domain.ItemType_CONSUMABLE_WEAPON
	case server.Armor:
		return domain.ItemType_ARMOR
	case server.Consumable:
		return domain.ItemType_CONSUMABLE
	case server.Resource:
		return domain.ItemType_RESOURCE
	case server.Quest:
		return domain.ItemType_QUEST_ITEM
	default:
		return domain.ItemType_RESOURCE
	}
}

func ToDTOItemType(itemType domain.ItemType) server.ItemType {
	switch itemType {
	case domain.ItemType_MELEE_WEAPON:
		return server.MeleeWeapon
	case domain.ItemType_RANGED_WEAPON:
		return server.RangedWeapon
	case domain.ItemType_CONSUMABLE_WEAPON:
		return server.ConsumableWeapon
	case domain.ItemType_ARMOR:
		return server.Armor
	case domain.ItemType_CONSUMABLE:
		return server.Consumable
	case domain.ItemType_RESOURCE:
		return server.Resource
	case domain.ItemType_QUEST_ITEM:
		return server.Quest
	default:
		return server.Resource
	}
}

// === MAP DTO TO DOMAIN MODEL ===
func ToItem(item *server.Item) *domain.Item {
	domainItemType := ToItemType(item.Type)
	domainItem := &domain.Item{
		ItemMeta: &domain.ItemMeta{
			Id: item.Id,
			Shape: &domain.Shape{
				RawShape: item.Shape.Rawshape,
				Height:   item.Shape.Height,
				Width:    item.Shape.Width,
			},
			Weight:   item.Weight,
			MaxStack: item.MaxStack,
		},
		Name:       item.Name,
		Text:       item.Text,
		Type:       domainItemType,
		Variant:    item.Variant,
		Durability: item.Durability,
		SellValue:  item.SellValue,
		BuyValue:   item.BuyValue,
	}
	return domainItem
}

func ToInventoryItem(inventoryItem *server.InventoryItem) *domain.InventoryItem {
	domainItem := ToItem(&inventoryItem.Item)
	domainInventoryItem := &domain.InventoryItem{
		Item: domainItem,
		Position: &domain.Position{
			X:        inventoryItem.Position.X,
			Y:        inventoryItem.Position.Y,
			Rotation: inventoryItem.Position.Rotation,
		},
		Quantity:       inventoryItem.Quantity,
		DurabilityLeft: inventoryItem.DurabilityLeft,
	}
	return domainInventoryItem
}

func ToInventoryMeta(inventoryMeta *server.InventoryMeta) *domain.InventoryMeta {
	domainInventoryMeta := &domain.InventoryMeta{
		Id:        inventoryMeta.Id,
		Name:      inventoryMeta.Name,
		UserId:    inventoryMeta.UserId,
		Width:     inventoryMeta.Volume.Width,
		Height:    inventoryMeta.Volume.Height,
		MaxWeight: inventoryMeta.MaxWeight,
	}
	return domainInventoryMeta
}

func ToInventory(inventory *server.Inventory) *domain.Inventory {
	items := make([]*domain.InventoryItem, 0)
	for _, item := range inventory.Items {
		items = append(items, ToInventoryItem(&item))
	}
	domainInventory := &domain.Inventory{
		InventoryMeta: ToInventoryMeta(&inventory.InventoryMeta),
		Items:         items,
	}
	return domainInventory
}

func ToUser(user *server.User) *domain.User {
	domainInventories := make([]*domain.InventoryMeta, 0)
	for _, inventory := range user.Inventories {
		domainInventories = append(domainInventories, ToInventoryMeta(&inventory))
	}
	domainUser := &domain.User{
		Id:          user.Id,
		Username:    user.Username,
		Inventories: domainInventories,
		Email:       user.Email,
	}
	return domainUser
}

// === MAP DOMAIN TO DTO MODEL ===
func ToDTOItem(item *domain.Item) server.Item {
	itemType := ToDTOItemType(item.Type)
	serverItem := &server.Item{
		BuyValue:   item.BuyValue,
		Durability: item.Durability,
		Id:         item.ItemMeta.Id,
		MaxStack:   item.ItemMeta.MaxStack,
		Name:       item.Name,
		SellValue:  item.SellValue,
		Shape: server.ItemShape{
			Height:   item.ItemMeta.Shape.Height,
			Rawshape: item.ItemMeta.Shape.RawShape,
			Width:    item.ItemMeta.Shape.Width,
		},
		Text:    item.Text,
		Type:    itemType,
		Variant: item.Variant,
		Weight:  item.ItemMeta.Weight,
	}
	return *serverItem
}

func ToDTOInventoryItem(inventoryItem *domain.InventoryItem) server.InventoryItem {
	serverInventoryItem := &server.InventoryItem{
		DurabilityLeft: inventoryItem.DurabilityLeft,
		Item:           ToDTOItem(inventoryItem.Item),
		Position: server.Position{
			Rotation: inventoryItem.Position.Rotation,
			X:        inventoryItem.Position.X,
			Y:        inventoryItem.Position.Y,
		},
		Quantity: inventoryItem.Quantity,
	}
	return *serverInventoryItem
}

func ToDTOInventoryMeta(inventoryMeta *domain.InventoryMeta) server.InventoryMeta {
	serverInventoryMeta := &server.InventoryMeta{
		Id:     inventoryMeta.Id,
		Name:   inventoryMeta.Name,
		UserId: inventoryMeta.UserId,
		Volume: server.Volume{
			Height: inventoryMeta.Height,
			Width:  inventoryMeta.Width,
		},
		MaxWeight: inventoryMeta.MaxWeight,
	}
	return *serverInventoryMeta
}

func ToDTOInventory(inventory *domain.Inventory) server.Inventory {
	serverInventoryItems := make([]server.InventoryItem, 0)
	for _, item := range inventory.Items {
		serverInventoryItems = append(serverInventoryItems, ToDTOInventoryItem(item))
	}
	serverInventory := &server.Inventory{
		InventoryMeta: ToDTOInventoryMeta(inventory.InventoryMeta),
		Items:         serverInventoryItems,
	}
	return *serverInventory
}

func ToDTOUser(user *domain.User) *server.User {
	serverInventories := make([]server.InventoryMeta, 0)
	for _, inventory := range user.Inventories {
		serverInventories = append(serverInventories, ToDTOInventoryMeta(inventory))
	}
	serverUser := &server.User{
		Email:       user.Email,
		Id:          user.Id,
		Inventories: serverInventories,
		Username:    user.Username,
	}
	return serverUser
}
