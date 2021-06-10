package ula

import (
	`github.com/storezhang/gox`
)

// 创建和直播返回
type andLiveCreateRsp struct {
	andLiveBaseRsp

	// 创建活动编号
	Id int64 `json:"id"`
	// 咪咕直播频道
	MiguChannelId string `json:"miguChannelId"`
	// 直播名称 最大长度为128个字
	Name string `json:"name"`
	// Status 状态
	// 0：未开始
	// 1：正在直播
	// 2：直播结束
	Status int64 `json:"status"`
	// 播放地址
	Urls gox.StringSlice `json:"urls"`
	// 分享地址
	ShareUrl string `json:"shareurl"`
	// 创建人编号
	UserId int64 `json:"userid"`
	// 创建人名字
	Username string `json:"username"`
	// 直播开始时间
	StartTime gox.Timestamp `json:"startTime"`
	// 直播结束时间,结束时间必须大于开始时间
	EndTime gox.Timestamp `json:"endTime"`
	// 推流地址
	PushUrl string `json:"pushurl"`
	// 是否允许播放录像
	// 0：否
	// 1：是
	EnabledPlay int64 `json:"enabledPlay"`
	// 是否录像
	// 0: 否
	// 1: 是
	SaveVideo int64 `json:"saveVideo"`
}
