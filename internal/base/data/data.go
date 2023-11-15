package data

import (
	"context"
	"github.com/baoyxing/geo-rpc/internal/base/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

type Data struct {
	Log      klog.CtxLogger
	Conf     *conf.Config
	Searcher *xdb.Searcher
}

func NewData(c *conf.Config, log klog.CtxLogger) (*Data, func(), error) {
	cBuff, err := xdb.LoadContentFromFile(c.Ip2RegionPath)
	ctx := context.Background()
	if err != nil {
		log.CtxFatalf(ctx, "xdb LoadContentFromFile failure,err:%s", err.Error())
	}
	searcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		log.CtxFatalf(ctx, "xdb NewWithBuffer failure,err:%s", err.Error())
	}
	d := &Data{
		Log:      log,
		Conf:     c,
		Searcher: searcher,
	}
	return d, func() {
	}, nil
}
