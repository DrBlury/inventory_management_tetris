package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func InventoryItemDTOFromDomain(i *domain.InventoryItem) server.InventoryItem {
	return server.InventoryItem{
		DurabilityLeft: i.DurabilityLeft,
		// TODO check what to do here
		//Item:     ItemDTOfromDomain(),
		Position: DTOPositionFromDomain(&i.Position),
		Quantity: i.Quantity,
	}
}

func InventoryItemDomainFromDTO(i *server.InventoryItem) domain.InventoryItem {
	return domain.InventoryItem{
		DurabilityLeft: i.DurabilityLeft,
		// TODO check what to do here
		Position: DomainPositionFromDTO(&i.Position),
		Quantity: i.Quantity,
	}
}
