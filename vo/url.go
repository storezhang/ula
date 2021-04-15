package vo

// Url 描述一个
type Url struct {
	// Type 视频格式
	Type FormatType `json:"type" yaml:"type" validate:"required,oneof=hls flv mp4 rtmp rtc"`
	// Link 视频链接地址
	Link string `json:"link" yaml:"link" validate:"required,url"`
}
