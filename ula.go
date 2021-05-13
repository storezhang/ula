package ula

import (
	"github.com/go-resty/resty/v2"
	"github.com/storezhang/ula/chuangcache"
	"github.com/storezhang/ula/migu"
	"github.com/storezhang/ula/vo"

	"github.com/storezhang/ula/and"
	"github.com/storezhang/ula/tencentyun"
)

// Live 直播接口
type Live interface {
	// Create 创建直播
	Create(create vo.Create) (id string, err error)
	// GetPushUrls 获得推流地址
	GetPushUrls(id string) (urls []vo.Url, err error)
	// GetPullCameras 获得拉流地址
	GetPullCameras(id string) (cameras []vo.Camera, err error)
}

// NewLive 创建通用的直播实现
func NewLive(config Config, resty *resty.Request) (live Live) {
	switch config.Type {
	case TypeAndLive:
		live = and.NewLive(config.And, resty)
	case TypeTencentyun:
		live = tencentyun.NewLive(config.Tencentyun)
	case TypeChuangcache:
		live = chuangcache.NewLive(config.Chuangcache)
	case TypeMigu:
		live = migu.NewLive(config.Migu, resty)
	}

	return
}
