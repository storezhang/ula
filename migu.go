package ula

import (
	`encoding/json`
	`fmt`
	`net/url`
	`sort`
	`strconv`
	`strings`
	`time`

	`github.com/storezhang/gox`
)

var _ executor = (*migu)(nil)

// 咪咕直播
type migu struct{}

func (m *migu) createLive(req *CreateLiveReq, options *options) (id string, err error) {
	createReq := &miguCreateReq{
		Title:     req.Title,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Subject:   req.Title,
	}
	createRsp := new(miguCreateRsp)
	if err = m.invoke(m.createEndpoint(options), createReq, createRsp, gox.HttpMethodPost, options); nil != err {
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
	if err = m.invoke(m.pushEndpoint(options), pushReq, pushRsp, gox.HttpMethodGet, options); nil != err {
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
	if err = m.invoke(m.pullEndpoint(options), pullReq, pullRsp, gox.HttpMethodGet, options); nil != err {
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
	stopReq := &miguForbidReq{
		ChannelId: id,
	}
	stopRsp := new(miguBaseRsp)
	if err = m.invoke(m.stopEndpoint(options), stopReq, stopRsp, gox.HttpMethodGet, options); nil != err {
		return
	}
	success = 0 == stopRsp.Ret

	return
}

func (m *migu) createEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/live/createChannel", options.migu.endpoint)
}

func (m *migu) pushEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/addr/getPushUrl", options.migu.endpoint)
}

func (m *migu) pullEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/addr/getPullUrl", options.migu.endpoint)
}

func (m *migu) stopEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/ctrl/forbidChannel", options.migu.endpoint)
}

func (m *migu) invoke(api string, req interface{}, rsp interface{}, method gox.HttpMethod, options *options) (err error) {
	currentTimeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	// 指定12小时后过期
	expired := strconv.FormatInt(time.Now().Add(12*time.Hour).UnixNano()/1e6, 10)
	random := gox.RandDigit(8)
	apiId := "001"
	grantType := "v2.0"
	encryptParams := map[string]string{
		"uid":              options.migu.uid,
		"secretId":         options.migu.secretId,
		"currentTimeStamp": currentTimeStamp,
		"expired":          expired,
		"random":           random,
		"apiId":            apiId,
		"grant_type":       grantType,
	}

	// 区分不同的方法
	// Get方法，body值是%5B%5D
	// Post方法，body值是JSON字符串编码后的值
	var getReqMap map[string]string
	if gox.HttpMethodGet == method {
		var reqMap map[string]interface{}
		if reqMap, err = gox.StructToMap(req); nil != err {
			return
		}

		getReqMap = make(map[string]string)
		for reqKey, reqValue := range reqMap {
			encryptParams[reqKey] = fmt.Sprintf("%v", reqValue)
			getReqMap[reqKey] = fmt.Sprintf("%v", reqValue)
		}
		encryptParams["body"] = "%5B%5D"
	} else {
		var reqJson []byte
		if reqJson, err = json.Marshal(req); nil != err {
			return
		}

		encryptParams["body"] = strings.Replace(url.QueryEscape(string(reqJson)), "+", "%20", -1)
	}

	// 请求参数按键排序
	var keys []string
	for key := range encryptParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	// 组装成地址格式，形如key=value&key=value
	params := make([]string, 0, len(encryptParams))
	for _, key := range keys {
		params = append(params, fmt.Sprintf("%s=%s", key, encryptParams[key]))
	}
	encryptData := strings.Join(params, "&")

	// 加密生成Token
	var token string
	if token, err = gox.Sha256Hmac(encryptData, options.migu.secretKey); nil != err {
		return
	}

	// 发出真正的请求
	api = fmt.Sprintf(
		"%s?apiId=%s&currentTimeStamp=%s&expired=%s&grant_type=%s&random=%s&secretId=%s&uid=%s&token=%s",
		api,
		apiId,
		currentTimeStamp,
		expired,
		grantType,
		random,
		options.migu.secretId,
		options.migu.uid,
		token,
	)
	if gox.HttpMethodGet == method {
		_, err = options.req().SetQueryParams(getReqMap).SetResult(rsp).Get(api)
	} else {
		_, err = options.req().SetBody(req).SetResult(rsp).Post(api)
	}

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
