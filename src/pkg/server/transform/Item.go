package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func ItemDTOfromDomain(i *domain.Item) server.Item {
	dtoShape := DTOShapeFromDomain(&i.Shape)
	dtoItemType := DTOItemTypeFromDomain(&i.Type)
	return server.Item{
		BuyValue:    i.BuyValue,
		Description: i.Description,
		Durability:  i.Durability,
		Id:          i.ID,
		MaxStack:    i.MaxStack,
		Name:        i.Name,
		SellValue:   i.SellValue,
		Shape:       dtoShape,
		Type:        dtoItemType,
		Variant:     i.Variant,
		Weight:      i.Weight,
	}
}

func ItemDomainFromDTO(i *server.Item) domain.Item {
	domainShape := DomainShapeFromDTO(&i.Shape)
	domainItemType := DomainItemTypeFromDTO(&i.Type)
	return domain.Item{
		BuyValue:    i.BuyValue,
		Description: i.Description,
		Durability:  i.Durability,
		ID:          i.Id,
		MaxStack:    i.MaxStack,
		Name:        i.Name,
		SellValue:   i.SellValue,
		Shape:       domainShape,
		Type:        domainItemType,
		Variant:     i.Variant,
		Weight:      i.Weight,
	}
}
