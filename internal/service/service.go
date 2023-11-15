package service

import (
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Service struct {
	log     klog.CtxLogger
	conf    *conf.Config
	geoRepo GeoRepo
}

func NewService(geoRepo GeoRepo,
	c *conf.Config, log klog.CtxLogger) geo.GeoService {
	return &Service{
		geoRepo: geoRepo,
		log:     log,
		conf:    c,
	}
}
