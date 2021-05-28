package ula

import (
	`time`

	`github.com/storezhang/gox`
)

type options struct {
	// 过期时间
	expired time.Duration
	// 协议
	scheme gox.URIScheme

	// 和直播
	andLive andLiveConfig
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
		andLive: andLiveConfig{
			endpoint: "http://dbtadmin.heshangwu.migucloud.com",
		},

		ulaType: TypeTencentyun,
	}
}
