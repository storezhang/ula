package ula

import (
	"github.com/storezhang/ula/conf"
)

// Config 配置
type Config struct {
	// Type 类型
	Type Type `json:"type" yaml:"type" validate:"required,oneof=and tencentyun"`
	// And 和直播配置
	And conf.AndLive `json:"and" yaml:"and" validate:"structonly"`
	// Tencentyun 腾讯云直播配置
	Tencentyun conf.Tencentyun `json:"tencentyun" yaml:"tencentyun" validate:"structonly"`
	// Chuangcache 创世云配置
	Chuangcache conf.Chuangcache `json:"chuangcache" yaml:"chuangcache" validate:"structonly"`
	// migu 咪咕配置
	Migu conf.Migu `json:"migu" yaml:"migu" validate:"structonly"`
}
