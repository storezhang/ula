package ula

import (
	"github.com/storezhang/ula/conf"
)

// Config 配置
type Config struct {
	// Type 类型
	Type Type `json:"type" yaml:"type" validate:"required,oneof=and tencentyunLive"`
	// And 和直播配置
	And conf.AndLive `json:"and" yaml:"and" validate:"structonly"`
	// Tencentyun 腾讯云直播配置
	Tencentyun conf.Tencentyun `json:"tencentyunLive" yaml:"tencentyunLive" validate:"structonly"`
	// Chuangcache 创世云配置
	Chuangcache conf.Chuangcache `json:"chuangcache" yaml:"chuangcache" validate:"structonly"`
}
