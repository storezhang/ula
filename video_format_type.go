package ula

const (
	// VideoFormatTypeHls HLS视频格式
	VideoFormatTypeHls VideoFormatType = "hls"
	// VideoFormatTypeFlv Flv视频格式
	VideoFormatTypeFlv VideoFormatType = "flv"
	// VideoFormatTypeMp4 Mp4视频格式
	VideoFormatTypeMp4 VideoFormatType = "mp4"
	// VideoFormatTypeRtmp Rtmp视频格式
	VideoFormatTypeRtmp VideoFormatType = "rtmp"
	// VideoFormatTypeRtc Webrtc视频格式
	VideoFormatTypeRtc VideoFormatType = "rtc"
)

// VideoFormatType 用于描述链接的格式
type VideoFormatType string
