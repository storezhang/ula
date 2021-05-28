package ula

import (
	"github.com/go-resty/resty/v2"
	"github.com/storezhang/gox"
)

type Ula interface {
	// CreateLive 创建直播信息
	CreateLive(req *CreateLiveReq, opts ...option) (id string, err error)
	// GetLivePushFlowInfo 获得推流信息
	GetLivePushFlowInfo(id string, opts ...option) (urls []Url, err error)
	// GetLivePullFlowInfo 获得拉流信息
	GetLivePullFlowInfo(id string, opts ...option) (cameras []Camera, err error)
}

// CreateLiveReq 创建一个直播
type CreateLiveReq struct {
	// Title 标题
	Title string `json:"title" yaml:"title" validate:"required"`
	// StartTime 开始时间
	StartTime gox.Timestamp `json:"startTime" yaml:"startTime"`
	// EndTime 结束时间
	EndTime gox.Timestamp `json:"endTime" yaml:"endTime"`
}

// New 创建适配器
func New(resty *resty.Request) Ula {
	return &ulaTemplate{
		andLive:        NewAndLive(resty),
		chuangcache:    NewChuangcache(),
		tencentyunLive: NewTencentyun(),
	}
}
