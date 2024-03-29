package client

import (
	"context"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo/geoservice"
	clientConf "github.com/baoyxing/micro-extend/pkg/config/client"
	"github.com/baoyxing/micro-extend/pkg/options/rpc"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/kitex-contrib/polaris"
	"time"
)

func NewRPCClientWithPolaris(ctx context.Context,
	confClient clientConf.Client, polaris clientConf.Polaris, jaeger clientConf.Jaeger,
	clientName, serverName string, suite *polaris.ClientSuite,
	log hlog.CtxLogger) (geoservice.Client, error) {
	opts, err := rpc.ClientOptions(confClient, polaris, jaeger, clientName, suite, log)
	if err != nil {
		log.CtxErrorf(ctx, "create node rpc client failure,err:%v", err)
		return nil, err
	}
	rpcClient, err := geoservice.NewClient(serverName, opts...)
	if err != nil {
		log.CtxErrorf(ctx, "%s 客户端连接失败 err：%s", serverName, err)
		return nil, err
	}
	log.CtxInfof(ctx, "%s 客户端连接成功 ", serverName)
	return rpcClient, nil
}

func NewRPCClient(ctx context.Context,
	conf clientConf.RpcClientConf, log hlog.CtxLogger) (geoservice.Client, error) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(conf.ServiceName),
		provider.WithExportEndpoint(conf.ProviderEndpoint),
		provider.WithInsecure(),
	)
	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {
			log.CtxFatalf(ctx, "provider Shutdown failure:%s", err.Error())
		}
	}(p, ctx)
	opts := make([]client.Option, 0)
	opts = append(opts, client.WithHostPorts(conf.Addr))
	opts = append(opts, client.WithMuxConnection(int(conf.MuxConnectionNum)))
	opts = append(opts, client.WithRPCTimeout(time.Duration(conf.RpcTimeout)*time.Second))
	opts = append(opts, client.WithSuite(tracing.NewClientSuite()))
	opts = append(opts, client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: conf.ServiceName}))
	rpcClient, err := geoservice.NewClient(conf.ServiceName, opts...)
	if err != nil {
		log.CtxErrorf(ctx, "%s 客户端连接失败 err：%s", conf.ServiceName, err)
		return nil, err
	}
	log.CtxInfof(ctx, "%s 客户端连接成功 ", conf.ServiceName)
	return rpcClient, nil
}
