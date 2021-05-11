package migu

type CameraPush struct {
	Status   int    `json:"status"`
	CamIndex string `json:"camIndex"`
	URL      string `json:"url"`
}

type PushUrl struct {
	UID        string       `json:"uid"`
	ChannelID  string       `json:"channelId"`
	ImgURL     string       `json:"imgUrl"`
	CameraList []CameraPush `json:"cameraList"`
}

type CameraPull struct {
	camIndex      string `json:"camIndex"`
	TranscodeList []struct {
		TransType string `json:"transType"`
		URLRtmp   string `json:"urlRtmp"`
		URLHls    string `json:"urlHls"`
		URLFlv    string `json:"urlFlv"`
	} `json:"transcodeList"`
}

type PullUrl struct {
	UID        string       `json:"uid"`
	ChannelID  string       `json:"channelId"`
	ImgURL     string       `json:"imgUrl"`
	CDNType    int8         `json:"cdnType"`
	ViewerNum  int64        `json:"viewerNum"`
	CameraList []CameraPull `json:"cameraList`
}
