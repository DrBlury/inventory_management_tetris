package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func DTOInventoryFromDomain(i *domain.Inventory) server.Inventory {
	serverInventoryItems := make([]server.InventoryItem, len(i.Items))
	for j, invItem := range i.Items {
		serverInventoryItems[j] = InventoryItemDTOFromDomain(&invItem)
	}

	return server.Inventory{
		Id:        i.InventoryMeta.ID,
		MaxWeight: i.InventoryMeta.MaxWeight,
		Name:      i.InventoryMeta.Name,
		UserId:    i.InventoryMeta.UserID,
		Volume: server.Volume{
			Width:  i.InventoryMeta.Width,
			Height: i.InventoryMeta.Height,
		},
		Items: serverInventoryItems,
	}
}

func DomainInventoryFromDTO(i *server.Inventory) domain.Inventory {
	domainInventoryItems := make([]domain.InventoryItem, len(i.Items))
	for j, invItem := range i.Items {
		domainInventoryItems[j] = InventoryItemDomainFromDTO(&invItem)
	}

	inventoryMeta := domain.InventoryMeta{
		ID:        i.Id,
		MaxWeight: i.MaxWeight,
		Name:      i.Name,
		UserID:    i.UserId,
		Width:     i.Volume.Width,
		Height:    i.Volume.Height,
	}

	return domain.Inventory{
		InventoryMeta: inventoryMeta,
		Items:         domainInventoryItems,
	}
}
