package repo

import (
	"context"
	"github.com/baoyxing/geo-rpc/internal/base/data"
	"github.com/baoyxing/geo-rpc/internal/service"
	"github.com/baoyxing/geo-rpc/kitex_gen/geo"
	"github.com/baoyxing/micro-extend/pkg/errors/rpc"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/pkg/errors"
	"strings"
)

type geoRepo struct {
	*data.Data
}

func NewGeoRepo(data *data.Data) service.GeoRepo {
	return &geoRepo{data}
}

func (repo *geoRepo) GetLocation(ctx context.Context) (*geo.Location, error) {
	ip, ok := metainfo.GetPersistentValue(ctx, "client_IP")
	if !ok {
		return nil, rpc.NewBizStatusError(rpc.ErrorTypeDataInvalid, errors.New("metainfo lack client_IP"))
	}
	loc, err := repo.Data.Searcher.SearchByStr(ip)
	if err != nil {
		return nil, rpc.NewBizStatusError(rpc.ErrorTypeDataInvalid, err)
	}
	locArr := strings.Split(loc, "|")
	country, region, city, isp := "unknown", "unknown", "unknown", "unknown"
	if len(locArr) == 5 {
		country = locArr[0]
		region = locArr[2]
		city = locArr[3]
		isp = locArr[4]
	}
	if country == "0" {
		country = "unknown"
	}
	if region == "0" {
		region = "unknown"
	} else {
		arr := strings.Split(region, "省")
		if len(arr) > 0 {
			region = arr[0]
		}
	}
	if city == "内网IP" || city == "0" {
		city = "unknown"
	} else {
		arr := strings.Split(city, "市")
		if len(arr) > 0 {
			city = arr[0]
		}
	}
	if isp == "内网IP" || isp == "0" {
		isp = "unknown"
	}
	return &geo.Location{
		Country: country,
		Region:  region,
		City:    city,
		Isp:     isp,
	}, nil

}
