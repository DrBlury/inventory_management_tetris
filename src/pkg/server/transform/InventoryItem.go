package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func InventoryItemDTOFromDomain(i *domain.InventoryItem) server.InventoryItem {
	return server.InventoryItem{
		DurabilityLeft: i.DurabilityLeft,
		Item:           ItemDTOfromDomain(&i.Item),
		Position:       DTOPositionFromDomain(&i.Position),
		Quantity:       i.Quantity,
	}
}

func InventoryItemDomainFromDTO(i *server.InventoryItem) domain.InventoryItem {
	return domain.InventoryItem{
		DurabilityLeft: i.DurabilityLeft,
		Item:           ItemDomainFromDTO(&i.Item),
		Position:       DomainPositionFromDTO(&i.Position),
		Quantity:       i.Quantity,
	}
}
