package ula

// Camera 描述一个直播摄像头
type Camera struct {
	// 摄像头编号
	Index int8 `json:"index" yaml:"index" validate:"required"`
	// 该摄像头对应的视频
	Videos []Video `json:"videos" yaml:"videos" validate:"structonly"`
}
