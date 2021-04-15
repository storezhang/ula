package vo

const (
	// VideoTypeAudio 音频
	VideoTypeAudio VideoType = "audio"
	// VideoType360P 音频
	VideoType360P VideoType = "360p"
	// VideoType480P 音频
	VideoType480P VideoType = "480p"
	// VideoType720P 音频
	VideoType720P VideoType = "720p"
	// VideoType1080P 音频
	VideoType1080P VideoType = "1080p"
	// VideoTypeOriginal 音频
	VideoTypeOriginal VideoType = "original"
)

// VideoType 视频类型
type VideoType string
