package ula

// Url 描述一个
type Url struct {
	// 视频格式
	Type VideoFormatType `json:"type" yaml:"type" validate:"required,oneof=hls flv mp4 rtmp rtc"`
	// 视频链接地址
	Link string `json:"link" yaml:"link" validate:"required,url"`
}
