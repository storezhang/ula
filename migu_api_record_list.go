package ula

type (
	miguListRecordReq struct {
		// 创建直播响应的channelId
		ChannelId string `json:"channelId"`
	}

	miguListRecordRsp struct {
		miguBaseRsp

		// 输出结果集
		Result struct {
			// 录制视频列表
			Contents []struct {
				// 录制文件地址
				RecordUrl string `json:"recordUrl"`
				// 录制状态
				// -1：录制中
				// 0：已录制
				// 1：导入中
				// 2：已导入
				Status int `json:"status"`
			} `json:"content"`
		} `json:"result"`
	}
)
