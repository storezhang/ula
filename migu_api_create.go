package ula

import (
	`github.com/storezhang/gox`
)

type (
	miguCreateReq struct {
		// 直播名称，最大长度为128个字
		Title string `json:"title"`
		// 直播开始时间
		StartTime gox.Timestamp `json:"startTime"`
		// 直播结束时间，直播结束时间，结束时间必须比开始时间大
		EndTime gox.Timestamp `json:"endTime"`
		// 直播主题，最大长度为128个字
		Subject string `json:"subject"`
		// 录制方式
		// 0：不录制
		// 2：录制
		Record int `json:"record"`
		// 是否导入云点播
		Demand int `json:"demand"`
		// 是否在导入点播后自动进行离线转码
		Transcode int `json:"transcode"`
	}

	miguCreateRsp struct {
		miguBaseRsp

		// 输出结果集，请求成功时会返回如下结果集，结 果集为空或请求异常时会返回null
		Result struct {
			// 直播ID
			ChannelId string `json:"channelId"`
		} `json:"result"`
	}
)
