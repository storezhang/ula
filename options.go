package ula

import (
	`crypto/tls`
	`time`

	`github.com/go-resty/resty/v2`
)

var defaultOptions = &options{
	resty:   resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}),
	expired: 3 * 24 * time.Hour,

	and: andConfig{
		endpoint: "http://dbtadmin.heshangwu.migucloud.com",
	},
	migu: miguConfig{
		endpoint: "https://api.migucloud.com",
	},
}

type options struct {
	// Http客户端
	resty *resty.Client

	// 过期时间
	expired time.Duration

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

func (o *options) req() *resty.Request {
	return o.resty.R()
}
