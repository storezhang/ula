package ula

// Video 视频
type Video struct {
	// Type 视频类型
	Type VideoType `json:"type" yaml:"type" validate:"required,oneof=audio 360p 480p 720p 1080p original"`
	// Urls 视频链接地址（可以是推流地址也可以是拉流地址）
	Urls []Url `json:"urls" yaml:"urls" validate:"structonly"`
}
