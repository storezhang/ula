package and

type baseAndLiveRsp struct {
	// ErrCode 错误码
	ErrCode int64 `json:"errcode"`
	// ErrMsg 错误信息
	ErrMsg string `json:"errmsg"`
	// ErrorDescription 错误描述
	ErrorDescription string `json:"error_description"`
	// Error 错误的字符串形式
	Error string `json:"error"`
}
