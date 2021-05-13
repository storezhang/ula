package vo

import (
	`github.com/storezhang/gox`
)

// Create 创建一个直播
type Create struct {
	// Title 标题
	Title string `json:"title" yaml:"title" validate:"required"`
	// StartTime 开始时间
	StartTime gox.Timestamp `json:"startTime" yaml:"startTime"`
	// EndTime 结束时间
	EndTime gox.Timestamp `json:"endTime" yaml:"endTime"`
	// 额外附加参数
	Extra map[string]string `json:"extra" yaml:"extra"`
}
