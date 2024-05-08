package apihandler

import (
	domain "linuxcode/inventory_manager/pkg/domain/model"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func MapEnumItemDTOfromDomain(dt domain.ItemType) server.ItemType {
	switch dt {
	case domain.ItemType_ARMOR:
		return server.Armor
	case domain.ItemType_CONSUMABLE:
		return server.Consumable
	case domain.ItemType_RANGED_WEAPON:
		return server.RangedWeapon
	case domain.ItemType_MELEE_WEAPON:
		return server.MeleeWeapon
	case domain.ItemType_CONSUMABLE_WEAPON:
		return server.ConsumableWeapon
	case domain.ItemType_QUEST_ITEM:
		return server.Quest
	case domain.ItemType_RESOURCE:
		return server.Resource
	default:
		return server.Consumable
	}
}

func MapEnumItemTypeToDomain(dt server.ItemType) domain.ItemType {
	switch dt {
	case server.Armor:
		return domain.ItemType_ARMOR
	case server.Consumable:
		return domain.ItemType_CONSUMABLE
	case server.RangedWeapon:
		return domain.ItemType_RANGED_WEAPON
	case server.MeleeWeapon:
		return domain.ItemType_MELEE_WEAPON
	case server.ConsumableWeapon:
		return domain.ItemType_CONSUMABLE_WEAPON
	case server.Quest:
		return domain.ItemType_QUEST_ITEM
	case server.Resource:
		return domain.ItemType_RESOURCE
	default:
		return domain.ItemType_CONSUMABLE
	}
}
