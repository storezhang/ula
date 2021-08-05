package ula

type andApiStopRsp struct {
	andApiBaseRsp

	// 直播编号
	Id int `json:"id"`
	// 是否成功
	Success bool `json:"success"`
}
