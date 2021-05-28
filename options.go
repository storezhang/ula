package ula

import (
	"time"

	"github.com/storezhang/gox"
)

type options struct {
	// 通信端点
	endpoint string
	// 授权密钥
	secret gox.Secret
	// 过期时间
	expired time.Duration

	chuangcache chuangcacheConfig

	livePull optionDomain
	livePush optionDomain

	// 类型
	ulaType Type
}

func defaultOptions() *options {
	return &options{
		chuangcache: chuangcacheConfig{},

		ulaType: TypeTencentyun,
	}
}
