package vo

const (
	// FormatTypeHls HLS视频格式
	FormatTypeHls FormatType = "hls"
	// FormatTypeFlv Flv视频格式
	FormatTypeFlv FormatType = "flv"
	// FormatTypeMp4 Mp4视频格式
	FormatTypeMp4 FormatType = "mp4"
	// FormatTypeRtmp Rtmp视频格式
	FormatTypeRtmp FormatType = "rtmp"
	// FormatTypeRtc Webrtc视频格式
	FormatTypeRtc FormatType = "rtc"
)

// FormatType 用于描述链接的格式
type FormatType string
