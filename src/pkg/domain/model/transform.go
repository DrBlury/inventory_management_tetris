package domain

import repo "linuxcode/inventory_manager/pkg/repo/generated"

func MapRepoInventoryToDomainInventoryMeta(rInv *repo.Inventory) *InventoryMeta {
	return &InventoryMeta{
		Id:        int64(rInv.ID),
		UserId:    int64(rInv.UserID.Int32),
		Name:      rInv.Invname.String,
		Width:     int64(rInv.Width.Int32),
		Height:    int64(rInv.Height.Int32),
		MaxWeight: int64(rInv.MaxWeight.Int32),
	}
}

// MapRepoInventoryItemsToDomainInventoryItems maps a slice
// of repo.InventoryItem to a slice of domain.InventoryItem
// ! It does NOT map the ItemMeta fields of the InventoryItem (only the ID is mapped)
func MapRepoInventoryItemsToDomainInventoryItems(rItems *[]repo.InventoryItem) []*InventoryItem {
	var domainInventoryItems = make([]*InventoryItem, 0, len(*rItems))
	for i, rItem := range *rItems {
		domainInventoryItems[i] = &InventoryItem{
			Item: &Item{
				ItemMeta: &ItemMeta{
					Id: int64(rItem.ItemID.Int32),
				},
			},
			Position: &Position{
				X:        int64(rItem.PositionX.Int32),
				Y:        int64(rItem.PositionY.Int32),
				Rotation: int64(rItem.Rotation.Int32),
			},
			Quantity:       int64(rItem.Quantity.Int32),
			DurabilityLeft: int64(rItem.DurabilityLeft.Int32),
		}
	}

	return domainInventoryItems
}

func MapEnumToProtobufEnum(itemType repo.ItemType) ItemType {
	switch itemType {
	case "consumable":
		return ItemType_CONSUMABLE
	case "armor":
		return ItemType_ARMOR
	case "rangedWeapon":
		return ItemType_RANGED_WEAPON
	case "meleeWeapon":
		return ItemType_MELEE_WEAPON
	case "consumableWeapon":
		return ItemType_CONSUMABLE_WEAPON
	case "quest":
		return ItemType_QUEST_ITEM
	case "resource":
		return ItemType_RESOURCE
	default:
		return ItemType_CONSUMABLE
	}
}

func MapRepoItemsToDomainItems(rItems ...repo.Item) []*Item {
	var domainItems = make([]*Item, 0, len(rItems))
	for _, rItem := range rItems {
		itemType := MapEnumToProtobufEnum(rItem.Type.ItemType)
		domainItems = append(domainItems, &Item{
			ItemMeta: &ItemMeta{
				Id: int64(rItem.ID),
				Shape: &Shape{
					Width:    int64(rItem.Width.Int32),
					Height:   int64(rItem.Height.Int32),
					RawShape: rItem.Rawshape.String,
				},
				Weight:   int64(rItem.Weight.Int32),
				MaxStack: int64(rItem.MaxStack.Int32),
			},
			Name:       rItem.Name.String,
			Text:       rItem.Text.String,
			Type:       itemType,
			Durability: int64(rItem.Durability.Int32),
			BuyValue:   int64(rItem.BuyValue.Int32),
			SellValue:  int64(rItem.SellValue.Int32),
			Variant:    rItem.Variant.String,
		})
	}
	return domainItems
}
