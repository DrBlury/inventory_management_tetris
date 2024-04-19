package domain

import repo "linuxcode/inventory_manager/pkg/repo/generated"

func MapRepoInventoryToDomainInventoryMeta(rInv *repo.Inventory) *InventoryMeta {
	return &InventoryMeta{
		ID:        int(rInv.ID),
		UserID:    int(rInv.UserID.Int32),
		Name:      rInv.Invname.String,
		Width:     int(rInv.Width.Int32),
		Height:    int(rInv.Height.Int32),
		MaxWeight: int(rInv.MaxWeight.Int32),
	}
}

// MapRepoInventoryItemsToDomainInventoryItems maps a slice
// of repo.InventoryItem to a slice of domain.InventoryItem
// ! It does NOT map the ItemMeta fields of the InventoryItem (only the ID is mapped)
func MapRepoInventoryItemsToDomainInventoryItems(rItems *[]repo.InventoryItem) *[]InventoryItem {
	var domainInventoryItems = make([]InventoryItem, 0, len(*rItems))
	for i, rItem := range *rItems {
		domainInventoryItems[i] = InventoryItem{
			Item: Item{
				ItemMeta: ItemMeta{
					ID: int(rItem.ItemID.Int32),
				},
			},
			Position: Position{
				X:        int(rItem.PositionX.Int32),
				Y:        int(rItem.PositionY.Int32),
				Rotation: int(rItem.Rotation.Int32),
			},
			Quantity:       int(rItem.Quantity.Int32),
			DurabilityLeft: int(rItem.DurabilityLeft.Int32),
		}
	}

	return &domainInventoryItems
}

func MapRepoItemsToDomainItems(rItems ...repo.Item) *[]Item {
	var domainItems = make([]Item, 0, len(rItems))
	for _, rItem := range rItems {
		itemType := ItemType(rItem.Type.ItemType)
		domainItems = append(domainItems, Item{
			ItemMeta: ItemMeta{
				ID: int(rItem.ID),
				Shape: Shape{
					Width:    int(rItem.Width.Int32),
					Height:   int(rItem.Height.Int32),
					RawShape: rItem.Rawshape.String,
				},
				Weight:   int(rItem.Weight.Int32),
				MaxStack: int(rItem.MaxStack.Int32),
			},
			Name:       rItem.Name.String,
			Text:       rItem.Text.String,
			Type:       itemType,
			Durability: int(rItem.Durability.Int32),
			BuyValue:   int(rItem.BuyValue.Int32),
			SellValue:  int(rItem.SellValue.Int32),
			Variant:    rItem.Variant.String,
		})
	}
	return &domainItems
}
