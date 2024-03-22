package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func ItemDTOfromDomain(i *domain.Item) server.Item {
	dtoShape := DTOShapeFromDomain(&i.ItemMeta.Shape)
	dtoItemType := DTOItemTypeFromDomain(&i.Type)
	return server.Item{
		BuyValue:    i.BuyValue,
		Description: i.Description,
		Durability:  i.Durability,
		Id:          i.ItemMeta.ID,
		MaxStack:    i.ItemMeta.MaxStack,
		Name:        i.Name,
		SellValue:   i.SellValue,
		Shape:       dtoShape,
		Type:        dtoItemType,
		Variant:     i.Variant,
		Weight:      i.ItemMeta.Weight,
	}
}

func ItemMetaFromDTO(i *server.Item) domain.ItemMeta {
	return domain.ItemMeta{
		ID:       i.Id,
		Shape:    DomainShapeFromDTO(&i.Shape),
		Weight:   i.Weight,
		MaxStack: i.MaxStack,
	}
}

func ItemDomainFromDTO(i *server.Item) domain.Item {
	domainItemType := DomainItemTypeFromDTO(&i.Type)
	domainItemMeta := ItemMetaFromDTO(i)
	return domain.Item{
		ItemMeta:    domainItemMeta,
		BuyValue:    i.BuyValue,
		Description: i.Description,
		Durability:  i.Durability,
		Name:        i.Name,
		SellValue:   i.SellValue,
		Type:        domainItemType,
		Variant:     i.Variant,
	}
}
