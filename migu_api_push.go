package ula

type miguPushRsp struct {
	miguBaseRsp

	// 输出结果集，请求成功时会返回如下结果集，结 果集为空或请求异常时会返回null
	Result struct {
		// 直播编号
		ChannelId string `json:"channelId"`
		// 用户编号
		Uid string `json:"uid"`
		// 直播封面
		ImgUrl string `json:"imgUrl"`
		// 机位列表
		CameraList []struct {
			// 机位状态
			// 0：未开始
			// 1：直播中
			// 2：暂停
			// 3：结束
			Status int `json:"status"`
			// 机位编号
			CamIndex string `json:"camIndex"`
			// 推流地址
			Url string `json:"url"`
		} `json:"cameraList"`
	} `json:"result"`
}
