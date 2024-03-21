package server

import (
	"linuxcode/inventory_manager/pkg/domain"
	server "linuxcode/inventory_manager/pkg/server/generated"
)

func DTOShapeFromDomain(s *domain.Shape) server.ItemShape {
	return server.ItemShape{
		Rawshape: s.RawShape,
		Width:    s.Width,
		Height:   s.Height,
	}
}

func DomainShapeFromDTO(s *server.ItemShape) domain.Shape {
	return domain.Shape{
		RawShape: s.Rawshape,
		Width:    s.Width,
		Height:   s.Height,
	}
}
