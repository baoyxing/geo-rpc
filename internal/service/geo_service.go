package service

import (
	"context"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo"
)

type GeoRepo interface {
	GetLocation(ctx context.Context) (*geo.Location, error)
}

func (s *Service) GetLocation(ctx context.Context) (*geo.Location, error) {
	return s.geoRepo.GetLocation(ctx)
}
