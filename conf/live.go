package conf

import (
	"time"

	"github.com/storezhang/gox"
)

// Live 直播配置
type Live struct {
	// Push 推流域名
	Push Domain `json:"push" yaml:"push"`
	// Pull 推流域名
	Pull Domain `json:"pull" yaml:"pull"`
	// Expiration 过期时间
	Expiration time.Duration `default:"24h" json:"expiration" yaml:"expiration"`
	// Scheme 协议
	Scheme gox.URIScheme `default:"https" json:"scheme" yaml:"scheme"`
}
