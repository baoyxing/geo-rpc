package server

import (
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo/geoservice"
	"github.com/baoyxing/micro-extend/pkg/options/rpc"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

func NewRPCServer(s geo.GeoService, c *conf.Config, log klog.CtxLogger) server.Server {
	opts, _ := rpc.ServerOptions(c.Server, c.Service, log)
	return geoservice.NewServer(s, opts...)
}
