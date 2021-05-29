package ula

import (
	"github.com/storezhang/gox"
)

type (
	// getAndLiveEventReq 获取直播信息请求
	getAndLiveEventReq struct {
		// ClientId 客户端 id
		ClientId string `json:"client_id" validate:"required"`
		// AccessToken 访问token
		AccessToken string `json:"access_token" validate:"required"`
		// Id 创建活动编号
		Id int64 `json:"id" validate:"required"`
	}

	// getAndLiveEventRsp 获取和直播数据返回
	getAndLiveEventRsp struct {
		baseAndLiveRsp

		// Id 创建活动编号
		Id int64 `json:"id"`
		// MiguChannelId 咪咕直播频道
		MiguChannelId string `json:"miguChannelId"`
		// Name 直播名称 最大长度为128个字
		Name string `json:"name"`
		// Status 状态
		// 0：未开始
		// 1：正在直播
		// 2：直播结束
		Status int64 `json:"status"`
		// Urls 播放地址
		Urls gox.StringSlice `json:"urls"`
		// ShareUrl 分享地址
		ShareUrl string `json:"shareurl"`
		// UserId 创建人编号
		UserId int64 `json:"userid"`
		// Username 创建人名字
		Username string `json:"username"`
		// StartTime 直播开始时间
		StartTime gox.Timestamp `json:"startTime"`
		// EndTime 直播结束时间,结束时间必须大于开始时间
		EndTime gox.Timestamp `json:"endTime"`
		// PushUrl 推流地址
		PushUrl string `json:"pushurl"`
		// CdnType CDN 类型
		// 0：网宿CDN
		// 1：咪咕CDN
		// 2: Live CDN
		CdnType int64 `json:"cdnType,string"`
		// EnabledPlay 是否允许播放录像
		// 0：否
		// 1：是
		EnabledPlay int64 `json:"enabledPlay,string"`
		// SaveVideo 是否录像
		// 0: 否
		// 1: 是
		SaveVideo int64 `json:"saveVideo,string"`
		// AllowComment 是否开启评论
		// 0：不开启
		// 1：开启
		AllowComment int8 `json:"allowComment"`
	}
)
