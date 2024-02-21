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
		Id:        i.ID,
		MaxWeight: i.MaxWeight,
		Name:      i.Name,
		UserId:    i.UserID,
		Volume:    server.Volume{
			SizeH: i.Width,
			SizeV: i.Height,
		},
		Items: serverInventoryItems,
	}
}

func DomainInventoryFromDTO(i *server.Inventory) domain.Inventory {
	domainInventoryItems := make([]domain.InventoryItem, len(i.Items))
	for j, invItem := range i.Items {
		domainInventoryItems[j] = InventoryItemDomainFromDTO(&invItem)
	}

	return domain.Inventory{
		ID:        i.Id,
		MaxWeight: i.MaxWeight,
		Name:      i.Name,
		UserID:    i.UserId,
		Width:     i.Volume.SizeH,
		Height:    i.Volume.SizeV,
		Items:     domainInventoryItems,
	}
}
