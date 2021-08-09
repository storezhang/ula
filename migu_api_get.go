package ula

type (
	miguGetReq struct {
		// 创建直播响应的channelId
		ChannelId string `json:"id"`
	}

	miguGetRsp struct {
		miguBaseRsp

		// 输出结果集
		Result struct {
			// 直播开始时间
			StartTime int64 `json:"startTime"`
			// 直播结束时间
			EndTime int64 `json:"endTime"`
		} `json:"result"`
	}
)
