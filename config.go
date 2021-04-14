package ala

import (
	`github.com/storezhang/ala/conf`
)

// Config 配置
type Config struct {
	// Type 类型
	Type Type `json:"type" yaml:"type" validate:"required,oneof=and tencentyun"`
	// And 和直播配置
	And conf.AndLive `json:"and" yaml:"and" validate:"structonly"`
	// Tencentyun 腾讯云直播配置
	Tencentyun conf.Tencentyun `json:"tencentyun" yaml:"tencentyun" validate:"structonly"`
}
