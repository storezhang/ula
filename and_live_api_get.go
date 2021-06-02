package ula

import (
	`strings`
	`time`

	`github.com/storezhang/gox`
)

type (
	andLiveGetRsp struct {
		baseAndLiveRsp

		// 咪咕渠道
		MiguChannelId string `json:"miguChannelId"`
		// 推流地址
		PushUrl string `json:"pushurl"`
		// 云录制播放地址
		ShareUrls string `json:"shareurls"`
		// 拉流地址
		// 如果直播结束后，这个地址就是云录制播放地址
		Urls []string `json:"urls"`
		// 直播开始时间
		StartTime gox.Timestamp `json:"starttime"`
		// 直播结束时间
		EndTime gox.Timestamp `json:"endTime"`
	}

	andLiveGet struct {
		rsp  *andLiveGetRsp
		time time.Time
	}
)

func (g *andLiveGetRsp) miguId() string {
	return strings.Split(g.MiguChannelId, "_")[1]
}
