package ula

type (
	miguListVidReq struct {
		// 创建直播响应的channelId
		ChannelId string `json:"channelId"`
	}

	miguListVidRsp struct {
		miguBaseRsp

		// 输出结果集
		Result []string `json:"result"`
	}
)
