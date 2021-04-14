package ala

import (
	`github.com/class100/live-protocol-go`
	`github.com/go-resty/resty/v2`

	`github.com/storezhang/ala/and`
	`github.com/storezhang/ala/tencentyun`
)

// Live 直播接口
type Live interface {
	// Create 创建直播
	Create(req *protocol.CreateReq) (id string, err error)
	// GetPushUrl 获得推流地址
	GetPushUrl(id string) (rsp *protocol.GetPushRsp, err error)
	// GetPullUrl 获得拉流地址
	GetPullUrl(id string) (rsp *protocol.GetPullRsp, err error)
}

// NewLive 创建通用的直播实现
func NewLive(config Config, resty *resty.Request) (live Live) {
	switch config.Type {
	case TypeAndLive:
		live = and.NewLive(config.And, resty)
	case TypeTencentyun:
		live = tencentyun.NewLive(config.Tencentyun)
	}

	return
}
