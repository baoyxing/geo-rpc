// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/baoyxing/geo-rpc/internal/base/data"
	server2 "github.com/baoyxing/geo-rpc/internal/base/server"
	"github.com/baoyxing/geo-rpc/internal/repo"
	"github.com/baoyxing/geo-rpc/internal/service"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

// Injectors from wire.go:

// *polaris.Registry, *registry.Info
func initApp(ctxLogger klog.CtxLogger, config *conf.Config) (server.Server, func(), error) {
	dataData, cleanup, err := data.NewData(config, ctxLogger)
	if err != nil {
		return nil, nil, err
	}
	geoRepo := repo.NewGeoRepo(dataData)
	geoService := service.NewService(geoRepo, config, ctxLogger)
	serverServer := server2.NewRPCServer(geoService, config, ctxLogger)
	return serverServer, func() {
		cleanup()
	}, nil
}
