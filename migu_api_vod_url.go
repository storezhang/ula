package ula

type (
	miguVodUrlReq struct {
		// 视频编号
		Vid string `json:"vid"`
	}

	miguVodUrlRsp struct {
		miguBaseRsp

		// 输出结果集
		Result struct {
			// 视频上下线状态
			PublicFlag int `json:"publicFlag"`
			// 上下线状态描述
			Desc string `json:"desc"`
			// 标题
			Title string `json:"title"`
			// 集合参数项
			List []struct {
				// 码率类型
				VType string `json:"vtype"`
				// 地址
				VUrl string `json:"vurl"`
			} `json:"list"`
		} `json:"result"`
	}
)
