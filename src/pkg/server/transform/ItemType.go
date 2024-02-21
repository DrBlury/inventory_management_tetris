package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func DTOItemTypeFromDomain(i *domain.ItemType) server.ItemType {
	// map the domain item type enum to the dto item type enum
	return server.ItemType(*i)
}

func DomainItemTypeFromDTO(i *server.ItemType) domain.ItemType {
	// map the dto item type enum to the domain item type enum
	return domain.ItemType(*i)
}
