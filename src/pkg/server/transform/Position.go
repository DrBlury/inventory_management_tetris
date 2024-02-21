package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func DTOPositionFromDomain(p *domain.Position) server.Position {
	return server.Position{
		X:        p.X,
		Y:        p.Y,
		Rotation: p.Rotation,
	}
}

func DomainPositionFromDTO(p *server.Position) domain.Position {
	return domain.Position{
		X:        p.X,
		Y:        p.Y,
		Rotation: p.Rotation,
	}
}
