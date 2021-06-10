package ula

import (
	`time`

	`github.com/storezhang/gox`
)

type options struct {
	// 通信商战
	endpoint string

	// 过期时间
	expired time.Duration
	// 协议
	scheme gox.URIScheme

	// 和直播
	andLive andLiveConfig
	// 咪咕直播
	migu miguConfig
	// 腾讯云直播
	tencentyun tencentyunConfig

	// 拉流域名配置
	pullDomain optionDomain
	// 推流域名配置
	pushDomain optionDomain

	// 类型
	ulaType Type
}

func defaultOptions() *options {
	return &options{
		expired: 3 * 24 * time.Hour,
		scheme:  gox.URISchemeHttps,
	}
}
