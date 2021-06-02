package ula

type andLiveStopRsp struct {
	baseAndLiveRsp

	// 直播编号
	Id int `json:"id"`
	// 是否成功
	Success bool `json:"success"`
}
