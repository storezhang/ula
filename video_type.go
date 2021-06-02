package ula

const (
	// VideoTypeAudio 音频
	VideoTypeAudio VideoType = "audio"
	// VideoType360P 360P
	VideoType360P VideoType = "360p"
	// VideoType480P 480P
	VideoType480P VideoType = "480p"
	// VideoType720P 720P
	VideoType720P VideoType = "720p"
	// VideoType1080P 1080P
	VideoType1080P VideoType = "1080p"
	// VideoTypeOriginal 原画
	VideoTypeOriginal VideoType = "original"
)

// VideoType 视频类型
type VideoType string
