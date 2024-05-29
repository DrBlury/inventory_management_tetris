package dboTransform

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	repo "linuxcode/inventory_manager/pkg/repo/generated"
)

// === MAP REPO TO DOMAIN MODEL ===
func ToItem(item *repo.Item) *domain.Item {
	// TODO implement
	return nil
}

func ToInventoryItem(inventoryItem *repo.InventoryItem) *domain.InventoryItem {
	// TODO implement
	return nil
}

func ToInventory(inventory *repo.Inventory) *domain.Inventory {
	// TODO implement
	return nil
}

func ToUser(user *repo.User) *domain.User {
	// TODO implement
	return nil
}

func ToItemType(itemType repo.ItemType) domain.ItemType {
	switch itemType {
	case repo.ItemTypeArmor:
		return domain.ItemType_ARMOR
	case repo.ItemTypeConsumable:
		return domain.ItemType_CONSUMABLE
	case repo.ItemTypeResource:
		return domain.ItemType_RESOURCE
	case repo.ItemTypeConsumableWeapon:
		return domain.ItemType_CONSUMABLE_WEAPON
	case repo.ItemTypeMeleeWeapon:
		return domain.ItemType_MELEE_WEAPON
	case repo.ItemTypeRangedWeapon:
		return domain.ItemType_RANGED_WEAPON
	case repo.ItemTypeQuest:
		return domain.ItemType_QUEST_ITEM
	default:
		return domain.ItemType_RESOURCE
	}
}

func ToRepoItemType(itemType domain.ItemType) repo.ItemType {
	switch itemType {
	case domain.ItemType_ARMOR:
		return repo.ItemTypeArmor
	case domain.ItemType_CONSUMABLE:
		return repo.ItemTypeConsumable
	case domain.ItemType_RESOURCE:
		return repo.ItemTypeResource
	case domain.ItemType_CONSUMABLE_WEAPON:
		return repo.ItemTypeConsumableWeapon
	case domain.ItemType_MELEE_WEAPON:
		return repo.ItemTypeMeleeWeapon
	case domain.ItemType_RANGED_WEAPON:
		return repo.ItemTypeRangedWeapon
	case domain.ItemType_QUEST_ITEM:
		return repo.ItemTypeQuest
	default:
		return repo.ItemTypeResource
	}
}
