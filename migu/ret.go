package migu

type cameraPush struct {
	Status   int    `json:"status"`
	CamIndex string `json:"camIndex"`
	URL      string `json:"url"`
}

type pushUrl struct {
	UID        string       `json:"uid"`
	ChannelID  string       `json:"channelId"`
	ImgURL     string       `json:"imgUrl"`
	CameraList []cameraPush `json:"cameraList"`
}

type cameraPull struct {
	CamIndex      string `json:"camIndex"`
	TranscodeList []struct {
		TransType string `json:"transType"`
		URLRtmp   string `json:"urlRtmp"`
		URLHls    string `json:"urlHls"`
		URLFlv    string `json:"urlFlv"`
	} `json:"transcodeList"`
}

type pullUrl struct {
	UID        string       `json:"uid"`
	ChannelID  string       `json:"channelId"`
	ImgURL     string       `json:"imgUrl"`
	CDNType    int8         `json:"cdnType"`
	ViewerNum  int64        `json:"viewerNum"`
	CameraList []cameraPull `json:"cameraList`
}

type channelID struct {
	ChannelID string `json:"channelId"`
}

type ret struct {
	Ret  int         `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
