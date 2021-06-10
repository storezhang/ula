package ula

import (
	`encoding/base64`
	`encoding/json`
	`fmt`

	`github.com/forgoer/openssl`
	`github.com/go-resty/resty/v2`
)

var (
	_ Ula         = (*migu)(nil)
	_ ulaInternal = (*migu)(nil)
)

// 咪咕直播
type migu struct {
	resty    *resty.Request
	template ulaTemplate
}

// NewMigu 创建咪咕直播实现类
func NewMigu(resty *resty.Request) (live *migu) {
	live = &migu{
		resty: resty,
	}
	live.template = ulaTemplate{migu: live}

	return
}

func (m *migu) CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error) {
	return m.template.CreateLive(req, opts...)
}

func (m *migu) GetPushUrls(id string, opts ...Option) (urls []Url, err error) {
	return m.template.GetPushUrls(id, opts...)
}

func (m *migu) GetPullCameras(id string, opts ...Option) (cameras []Camera, err error) {
	return m.template.GetPullCameras(id, opts...)
}

func (m *migu) Stop(id string, opts ...Option) (success bool, err error) {
	return m.template.Stop(id, opts...)
}

func (m *migu) createLive(req *CreateLiveReq, options *options) (id string, err error) {
	createReq := &miguCreateReq{
		Title:     req.Title,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Subject:   req.Title,
	}
	createRsp := new(miguCreateRsp)
	if err = m.invoke(m.createEndpoint(options), createReq, createRsp, options); nil != err {
		return
	}
	id = createRsp.Result.ChannelId

	return
}

func (m *migu) getPushUrls(id string, options *options) (urls []Url, err error) {
	pushReq := &miguStreamReq{
		ChannelId: id,
	}
	pushRsp := new(miguPushRsp)
	if err = m.invoke(m.pushEndpoint(options), pushReq, pushRsp, options); nil != err {
		return
	}

	cameraCount := len(pushRsp.Result.CameraList)
	if 0 != cameraCount {
		urls = make([]Url, 0, cameraCount)
		for _, camera := range pushRsp.Result.CameraList {
			urls = append(urls, Url{
				Type: VideoFormatTypeRtmp,
				Link: camera.Url,
			})
		}
	}

	return
}

func (m *migu) getPullCameras(id string, options *options) (cameras []Camera, err error) {
	pullReq := &miguStreamReq{
		ChannelId: id,
	}
	pullRsp := new(miguPullRsp)
	if err = m.invoke(m.stopEndpoint(options), pullReq, pullRsp, options); nil != err {
		return
	}

	cameraCount := len(pullRsp.Result.CameraList)
	if 0 != cameraCount {
		cameras = make([]Camera, 0, cameraCount)
		for _, mc := range pullRsp.Result.CameraList { // mc是miguCamera的简写
			camera := Camera{}
			for _, transcode := range mc.TranscodeList {
				camera.Videos = append(camera.Videos, Video{
					Type: m.parseVideoType(transcode.TransType),
					Urls: []Url{{
						Type: VideoFormatTypeFlv,
						Link: transcode.UrlFlv,
					}, {
						Type: VideoFormatTypeHls,
						Link: transcode.UrlHls,
					}, {
						Type: VideoFormatTypeRtmp,
						Link: transcode.UrlRtmp,
					}},
				})
			}
			cameras = append(cameras, camera)
		}
	}

	return
}

func (m *migu) stop(id string, options *options) (success bool, err error) {
	stopReq := &miguStreamReq{
		ChannelId: id,
	}
	stopRsp := new(miguBaseRsp)
	if err = m.invoke(m.stopEndpoint(options), stopReq, stopRsp, options); nil != err {
		return
	}
	success = 0 == stopRsp.Ret

	return
}

func (m *migu) createEndpoint(options *options) string {
	return fmt.Sprintf("%s/eduOnlineApi/migu/createChannel", options.endpoint)
}

func (m *migu) pushEndpoint(options *options) string {
	return fmt.Sprintf("%s/eduOnlineApi/migu/getPushUrl", options.endpoint)
}

func (m *migu) stopEndpoint(options *options) string {
	return fmt.Sprintf("%s/eduOnlineApi/migu/forbidChannel", options.endpoint)
}

func (m *migu) invoke(url string, req interface{}, rsp interface{}, options *options) (err error) {
	var reqJson []byte
	if reqJson, err = json.Marshal(req); nil != err {
		return
	}

	key := []byte(options.migu.key)
	iv := []byte(options.migu.iv)

	var reqEncrypt []byte
	if reqEncrypt, err = openssl.AesCBCEncrypt(reqJson, key, iv, openssl.PKCS7_PADDING); nil != err {
		return
	}

	var miguRsp *resty.Response
	reqBody := map[string]string{
		"params": base64.StdEncoding.EncodeToString(reqEncrypt),
	}
	if miguRsp, err = m.resty.SetBody(reqBody).Post(fmt.Sprintf("%s?token=%s", url, options.migu.token)); nil != err {
		return
	}

	var rspDecrypt []byte
	if rspDecrypt, err = openssl.AesCBCDecrypt(miguRsp.Body(), key, iv, openssl.PKCS7_PADDING); nil != err {
		return
	}
	err = json.Unmarshal(rspDecrypt, rsp)

	return
}

func (m *migu) parseVideoType(transcodeType int) (videoType VideoType) {
	switch transcodeType {
	case 0:
		videoType = VideoTypeOriginal
	case 1:
		videoType = VideoType360P
	case 2:
		videoType = VideoType480P
	case 3:
		videoType = VideoType720P
	case 4:
		videoType = VideoType1080P
	case 5:
		videoType = VideoTypeAudio
	default:
		videoType = VideoTypeOriginal
	}

	return
}
