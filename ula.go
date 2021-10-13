package ula

import (
	`sync`

	`github.com/storezhang/gox`
)

type Ula interface {
	// CreateLive 创建直播信息
	CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error)
	// GetPushUrls 获得推流信息
	GetPushUrls(id string, opts ...Option) (urls []Url, err error)
	// GetPullCameras 获得拉流信息
	GetPullCameras(id string, opts ...Option) (cameras []Camera, err error)
	// Stop 结束直播
	Stop(id string, opts ...Option) (success bool, err error)
	// GetViewerNum 获取在线人数
	GetViewerNum(id string, opts ...Option) (viewerNum int64, err error)
}

// CreateLiveReq 创建一个直播
type CreateLiveReq struct {
	// 标题
	Title string `json:"title" yaml:"title" xml:"title" validate:"required"`
	// 开始时间
	StartTime gox.Timestamp `json:"startTime" yaml:"startTime" xml:"startTime"`
	// 结束时间
	EndTime gox.Timestamp `json:"endTime" yaml:"endTime" xml:"endTime"`
	// 相机个数
	Cameras int `json:"cameras" yaml:"cameras" xml:"cameras"`
}

// New 创建适配器
func New(opts ...Option) Ula {
	for _, opt := range opts {
		opt.apply(defaultOptions)
	}

	return &template{
		andLive: &and{
			tokenCache:  sync.Map{},
			getCache:    sync.Map{},
			recordCache: sync.Map{},
		},
		migu:        &migu{},
		tencentyun:  &tencentyun{},
		chuangcache: &chuangcache{},
	}
}
