package ula

type baseAndLiveRsp struct {
	// 错误码
	ErrCode int64 `json:"errcode"`
	// 错误信息
	ErrMsg string `json:"errmsg"`
	// 错误描述
	ErrorDescription string `json:"error_description"`
	// 错误的字符串形式
	Error string `json:"error"`
}
