package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func UserDTOFromDomain(u *domain.User) server.User {
	dtoInventories := make([]server.Inventory, len(u.Inventories))
	for i, inv := range u.Inventories {
		dtoInventories[i] = DTOInventoryFromDomain(&inv)
	}
	return server.User{
		Id:          u.ID,
		Inventories: dtoInventories,
		Username:    u.Username,
	}
}

func UserDomainFromDTO(u *server.User) domain.User {
	domainInventories := make([]domain.Inventory, len(u.Inventories))
	for i, inv := range u.Inventories {
		domainInventories[i] = DomainInventoryFromDTO(&inv)
	}
	return domain.User{
		ID:        u.Id,
		Inventories: domainInventories,
		Username:  u.Username,
	}
}
