package ula

import (
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

type Ula interface {
	// CreateLive 创建直播信息
	CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error)
	// GetPushUrls 获得推流信息
	GetPushUrls(id string, opts ...Option) (urls []Url, err error)
	// GetPullCameras 获得拉流信息
	GetPullCameras(id string, opts ...Option) (cameras []Camera, err error)
}

// CreateLiveReq 创建一个直播
type CreateLiveReq struct {
	// 标题
	Title string `json:"title" yaml:"title" validate:"required"`
	// 开始时间
	StartTime gox.Timestamp `json:"startTime" yaml:"startTime"`
	// 结束时间
	EndTime gox.Timestamp `json:"endTime" yaml:"endTime"`
}

// New 创建适配器
func New(resty *resty.Request) Ula {
	return &ulaTemplate{
		andLive:     NewAndLive(resty),
		tencentyun:  NewTencentyun(),
		chuangcache: NewChuangcache(),
	}
}
