package ula

import (
	`crypto/tls`
	`time`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

var defaultOptions = &options{
	resty:   resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
	expired: 3 * 24 * time.Hour,
	scheme:  gox.URISchemeHttps,

	and: andConfig{
		endpoint: "http://dbtadmin.heshangwu.migucloud.com",
	},
	migu: miguConfig{
		endpoint: "https://api.migucloud.com",
	},
}

type (
	// Option 选项
	Option interface {
		apply(options *options)
	}

	options struct {
		// Http客户端
		resty *resty.Client

		// 过期时间
		expired time.Duration
		// 生成的地址协议
		scheme gox.URIScheme

		// 和直播
		and andConfig
		// 咪咕直播
		migu miguConfig
		// 腾讯云直播
		tencentyun  tencentyunConfig
		chuangcache chuangcacheConfig

		// 类型
		ulaType Type
	}
)

// NewOptions 创建选项，因为option接口不对外暴露，如果用户想在外面创建option并赋值将无法完成，特意提供创建option的快捷方式
func NewOptions(opts ...Option) []Option {
	return opts
}

func (o *options) req() *resty.Request {
	return o.resty.R()
}
