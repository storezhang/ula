package and

import (
	"github.com/storezhang/gox"
)

type (
	createAndLiveEventReq struct {
		// ClientId 客户端 id
		ClientId string `json:"client_id" validate:"required"`
		// AccessToken 访问token
		AccessToken string `json:"access_token" validate:"required"`
		// Name 直播名称 最大长度为128个字
		Name string `json:"name" validate:"required,max=128"`
		// StartTime 直播开始时间
		StartTime gox.Timestamp `json:"startTime"`
		// EndTime 直播结束时间,结束时间必须大于开始时间
		EndTime gox.Timestamp `json:"endTime"`
		// 当前用户Id
		Uid int64 `json:"uid"`
	}

	// createAndLiveEventRsp 创建和直播返回
	createAndLiveEventRsp struct {
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
		// 0：分享地址
		ShareUrl string `json:"shareurl"`
		// UserId 创建人编号
		UserId int64 `json:"userid"`
		// Username 创建人名字
		Username string `json:"username"`
		// StartTime 直播开始时间
		StartTime gox.Timestamp `json:"startTime"`
		// EndTime 直播结束时间,结束时间必须大于开始时间
		EndTime gox.Timestamp `json:"endTime"`
		// pushurl 推流地址
		PushUrl string `json:"pushurl"`
		// EnabledPlay 是否允许播放录像
		// 0：否
		// 1：是
		EnabledPlay int64 `json:"enabledPlay"`
		// SaveVideo 是否录像
		// 0: 否
		// 1: 是
		SaveVideo int64 `json:"saveVideo"`
	}
)
