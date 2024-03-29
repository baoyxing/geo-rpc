package main

import (
	"context"
	"flag"
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/baoyxing/micro-extend/pkg/configuration/polaris"
	"github.com/baoyxing/micro-extend/pkg/utils/logutils"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexZap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"gopkg.in/yaml.v3"
	"time"
)

var (
	namespace = "edge"                    //空间名称
	fileGroup = "pcdn"                    //分组名称
	fileName  = "rpc/jobPlan/config.yaml" //文件名称
)

func main() {
	flag.Parse()
	klog.SetLogger(kitexZap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	configFile := polaris.ConfigApi(namespace, fileGroup, fileName)
	//解析远程配置文件
	config := new(conf.Config)
	//config = conf.GetConf()
	err := yaml.Unmarshal([]byte(configFile.GetContent()), &config)

	if err != nil {
		klog.CtxErrorf(context.Background(), "yaml 反序列化失败 error：%v", err)
		panic(err)
	}

	// 自定义日志配置
	if config.Logger.Enable {
		opts := make([]logutils.Option, 0, 8)
		opts = append(opts, logutils.WithPath(config.Logger.Path))
		opts = append(opts, logutils.WithMaxSize(config.Logger.MaxSize))
		opts = append(opts, logutils.WithMaxBackups(config.Logger.MaxBackups))
		opts = append(opts, logutils.WithMaxAge(config.Logger.MaxAge))
		opts = append(opts, logutils.WithCompress(config.Logger.Compress))
		opts = append(opts, logutils.WithOutputMode(config.Logger.OutputMode))
		opts = append(opts, logutils.WithRotationDuration(time.Duration(config.Logger.RotationDuration)))
		opts = append(opts, logutils.WithSuffix(config.Logger.Suffix))
		logutils.NewKitexLog(opts...)
		klog.SetLevel(logutils.Level(config.Logger.Level).KLogLevel())
	}
	//wire 依赖注入
	svr, cleanup, err := initApp(klog.DefaultLogger(), config)
	if err != nil {
		panic(err)
	}

	defer cleanup()
	if err := svr.Run(); err != nil {
		panic(err)
	}
}
