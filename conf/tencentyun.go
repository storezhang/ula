package conf

import (
	`time`

	`github.com/storezhang/gox`
)

type (
	// Tencentyun 腾讯云直播配置
	Tencentyun struct {
		// Key 鉴权密钥，主要用来生成防盗链地址
		Key Key `json:"key" yaml:"key"`
		// Domain 域名
		Domain Domain `json:"domain" yaml:"domain"`
		// Expiration 过期时间
		Expiration time.Duration `default:"24h" json:"expiration" yaml:"expiration"`
		// Scheme 协议
		Scheme gox.URIScheme `default:"https" json:"scheme" yaml:"scheme"`
	}

	// Key 鉴权密钥，主要用来生成防盗链地址
	Key struct {
		// Push 推流密钥
		Push string `json:"push" yaml:"push"`
		// Pull 拉流密钥
		Pull string `json:"pull" yaml:"pull"`
	}

	// Domain 域名
	Domain struct {
		// Push 推流域名
		Push string `json:"push" yaml:"push" validate:"required"`
		// Pull 拉流域名
		Pull string `json:"pull" yaml:"pull" validate:"required"`
		// BizId 支持400ms左右的超低延迟播放
		BizId int `json:"bizId" yaml:"bizId"`
	}
)
