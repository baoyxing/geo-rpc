//go:build wireinject
// +build wireinject

package main

import (
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/baoyxing/geo-rpc/internal/base/data"
	"github.com/baoyxing/geo-rpc/internal/base/server"
	"github.com/baoyxing/geo-rpc/internal/repo"
	"github.com/baoyxing/geo-rpc/internal/service"
	"github.com/cloudwego/kitex/pkg/klog"
	kserver "github.com/cloudwego/kitex/server"
	"github.com/google/wire"
)

// The build tag makes sure the stub is not built in the final build.

// *polaris.Registry, *registry.Info
func initApp(klog.CtxLogger, *conf.Config) (kserver.Server, func(), error) {
	panic(wire.Build(service.ProviderSet, repo.ProviderSet, server.ProviderSet, data.ProviderSet))
}
