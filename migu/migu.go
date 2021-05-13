package migu

import (
	`bytes`
	`encoding/json`
	`errors`
	`fmt`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/ula/conf`
	`github.com/storezhang/ula/vo`
)

type live struct {
	config conf.Migu
	resty  *resty.Request
}

// NewLive 创建咪咕直播实现类
func NewLive(config conf.Migu, resty *resty.Request) *live {
	return &live{
		config: config,
		resty:  resty,
	}
}

// generate token
func (l *live) genToken() string {
	return ""
}

func (l *live) Create(create vo.Create) (id string, err error) {
	if create.Extra == nil {
		return "", ParamMissing
	}
	if _, exist := create.Extra["subject"]; !exist {
		return "", ParamMissSubject
	}

	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/createChannel", l.config.Scheme, l.config.Addr)
	bodyMap := make(map[string]string, len(create.Extra)+3)
	for k, v := range create.Extra {
		bodyMap[k] = v
	}
	bodyMap["title"] = create.Title
	bodyMap["startTime"] = create.StartTime.Time().Format("2006-01-02 15:04:05")
	bodyMap["endTime"] = create.EndTime.Time().Format("2006-01-02 15:04:05")

	var bts []byte
	bts, err = json.Marshal(bodyMap)
	if err != nil {
		return
	}

	channelID := new(channelID)
	ret := &ret{
		Data: channelID,
	}
	_, err = l.resty.SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewBuffer(bts)).
		SetResult(ret).
		Post(url)
	if err != nil {
		return
	}

	if err = l.hasErr(ret); err != nil {
		return
	}

	return channelID.ChannelID, nil
}

func (l *live) GetPushUrls(channelID string) (urls []vo.Url, err error) {
	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/getPushUrl", l.config.Scheme, l.config.Addr)
	body := fmt.Sprintf(`{"channel_id": "%s"}`, channelID)

	result := new(pushUrl)
	ret := &ret{
		Data: result,
	}
	_, err = l.resty.SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewBufferString(body)).
		SetResult(ret).
		Post(url)
	if err != nil {
		return
	}

	if err = l.hasErr(ret); err != nil {
		return
	}

	urls = make([]vo.Url, len(result.CameraList))
	for i, camera := range result.CameraList {
		urls[i].Link = camera.URL
	}

	return
}

func (l *live) GetPullCameras(channelID string) (cameras []vo.Camera, err error) {
	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/getPullUrl", l.config.Scheme, l.config.Addr)
	body := fmt.Sprintf(`{"channel_id": "%s"}`, channelID)

	result := new(pullUrl)
	ret := &ret{
		Data: result,
	}
	_, err = l.resty.SetHeader("Content-Type", "application/json").
		SetBody(bytes.NewBufferString(body)).
		SetResult(ret).
		Post(url)
	if err != nil {
		return
	}

	if err = l.hasErr(ret); err != nil {
		return
	}

	cameras = make([]vo.Camera, 0, len(result.CameraList))
	for _, camera := range result.CameraList {

		videos := make([]vo.Video, 0, len(camera.TranscodeList))
		for _, v := range camera.TranscodeList {
			videos = append(videos, vo.Video{
				Type: transType2VideoType(v.TransType),
				Urls: []vo.Url{
					vo.Url{
						Type: vo.FormatTypeHls,
						Link: v.URLHls,
					},
					vo.Url{
						Type: vo.FormatTypeRtmp,
						Link: v.URLRtmp,
					},
					vo.Url{
						Type: vo.FormatTypeFlv,
						Link: v.URLFlv,
					},
				},
			})
		}

		cm := vo.Camera{
			Index:  camera.CamIndex,
			Videos: videos,
		}
		cameras = append(cameras, cm)
	}

	return

}

func (l *live) hasErr(ret *ret) (err error) {
	if ret.Ret != 0 {
		err = errors.New(ret.Msg)
		return
	}

	return
}

func transType2VideoType(transType string) vo.VideoType {
	switch transType {
	case "0":
		return vo.VideoTypeOriginal
	case "1":
		return vo.VideoType360P
	case "2":
		return vo.VideoType480P
	case "3":
		return vo.VideoType720P
	case "4":
		return vo.VideoType1080P
	case "5":
		return vo.VideoTypeAudio
	default:
		return vo.VideoType360P
	}
}
