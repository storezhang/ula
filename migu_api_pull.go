package ula

type miguPullRsp struct {
	miguBaseRsp

	// 输出结果集，请求成功时会返回如下结果集，结 果集为空或请求异常时会返回null
	Result struct {
		// 直播编号
		ChannelId string `json:"channelId"`
		// 用户编号
		Uid string `json:"uid"`
		// 直播封面
		ImgUrl string `json:"imgUrl"`
		// CDN类型
		// 0：低延时CDN
		// 1：在线直播CDN
		// 2：咪咕LiveCDN
		CdnType int `json:"cdnType"`
		// 观看人数
		ViewerNum int `json:"viewerNum"`
		// 机位列表
		CameraList []struct {
			// 机位编号
			CamIndex string `json:"camIndex"`
			// 转码列表
			TranscodeList []struct {
				// 转码编号
				TransIndex string `json:"transIndex"`
				// 转码类型（清晰度）
				// 与创建直播时，videoType中类型度)对应
				// 原画质：0
				// 流畅：1
				// 标清：2
				// 高清：3
				// 超清：4
				// 音频：5
				TransType int `json:"transType"`
				// FLV观察地址
				UrlFlv string `json:"urlFlv"`
				// HLS观察地址
				UrlHls string `json:"urlHls"`
				// RTMP观看地址
				UrlRtmp string `json:"urlRtmp"`
			} `json:"transcodeList"`
		} `json:"cameraList"`
	} `json:"result"`
}
