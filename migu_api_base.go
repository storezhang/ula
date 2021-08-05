package ula

type miguBaseRsp struct {
	// 状态码
	Ret int `json:"ret,string"`
	// 信息描述
	Msg string `json:"msg"`
}
