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
		CameraNum: 3,
		// 录制视频
		Record: 2,
		// 自动导入云点播（只有导入云点播后才能获取播放地址）
		Demand: 1,
		// 导入云点播后自动转码
		Transcode: 1,
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
	getReq := &miguGetReq{
		ChannelId: id,
	}
	getRsp := new(miguGetRsp)
	if err = m.invoke(m.getEndpoint(options), getReq, getRsp, gox.HttpMethodGet, options); nil != err {
		return
	}

	// 判断直播是否已经结束，如果结束，需要取得录制地址
	now := time.Now().UnixNano() / 1e6
	// 直播结束5分钟后会自动导入
	if now < (getRsp.Result.EndTime + int64(5*time.Minute/1e6)) {
		cameras, err = m.getPullUrls(id, options)
	} else {
		cameras, err = m.getRecordUrls(id, options)
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

func (m *migu) getPullUrls(id string, options *options) (cameras []Camera, err error) {
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
		for index, mc := range pullRsp.Result.CameraList { // mc是miguCamera的简写
			camera := Camera{
				Index: int8(index),
			}

			for _, transcode := range mc.TranscodeList {
				camera.Videos = append(camera.Videos, Video{
					Type: m.parseTranscodeType(transcode.TransType),
					Urls: []Url{{
						Type: VideoFormatTypeFlv,
						Link: m.parseFlv(transcode.UrlFlv, options.scheme),
					}, {
						Type: VideoFormatTypeHls,
						Link: m.parseHls(transcode.UrlHls, options.scheme),
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

func (m *migu) getRecordUrls(id string, options *options) (cameras []Camera, err error) {
	listReq := &miguListVidReq{
		ChannelId: id,
	}
	listRsp := new(miguListVidRsp)
	if err = m.invoke(m.listVidEndpoint(options), listReq, listRsp, gox.HttpMethodPost, options); nil != err {
		return
	}

	cameras = make([]Camera, 0, len(listRsp.Result))
	for index, vid := range listRsp.Result {
		urlReq := &miguVodUrlReq{
			Vid: vid,
		}
		urlRsp := new(miguVodUrlRsp)
		if urlErr := m.invoke(m.vodVerifyHttpsUrlEndpoint(options), urlReq, urlRsp, gox.HttpMethodGet, options); nil != urlErr {
			continue
		}

		camera := Camera{
			Index: int8(index),
		}
		for _, video := range urlRsp.Result.List {
			camera.Videos = append(camera.Videos, Video{
				Type: m.parseVType(video.VType),
				Urls: []Url{{
					Type: VideoFormatTypeHls,
					Link: video.VUrl,
				}},
			})
		}
		cameras = append(cameras, camera)
	}

	return
}

func (m *migu) createEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/live/createChannel", options.migu.endpoint)
}

func (m *migu) getEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/live/getChannel", options.migu.endpoint)
}

func (m *migu) pushEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/name/getPushUrl", options.migu.endpoint)
}

func (m *migu) pullEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/name/getPullUrl", options.migu.endpoint)
}

func (m *migu) stopEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/ctrl/forbidChannel", options.migu.endpoint)
}

func (m *migu) listRecordEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/record/listRecord", options.migu.endpoint)
}

func (m *migu) listVidEndpoint(options *options) string {
	return fmt.Sprintf("%s/l2/record/listVid", options.migu.endpoint)
}

func (m *migu) vodVerifyHttpsUrlEndpoint(options *options) string {
	return fmt.Sprintf("%s/vod2/v1/getUrlVerifyForHttps", options.migu.endpoint)
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
	// var miguRsp *resty.Response
	if gox.HttpMethodGet == method {
		_, err = options.req().SetQueryParams(getReqMap).SetResult(rsp).Get(api)
	} else {
		_, err = options.req().SetBody(req).SetResult(rsp).Post(api)
	}
	// fmt.Println(miguRsp)

	return
}

func (m *migu) parseTranscodeType(transcodeType int) (videoType VideoType) {
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

func (m *migu) parseVType(vType string) (videoType VideoType) {
	switch vType {
	case "流畅":
		videoType = VideoType360P
	case "标清":
		videoType = VideoType480P
	case "高清":
		videoType = VideoType720P
	case "超清":
		videoType = VideoType1080P
	case "原画质":
		videoType = VideoTypeOriginal
	default:
		videoType = VideoTypeOriginal
	}

	return
}

func (m *migu) parseFlv(original string, scheme gox.URIScheme) string {
	uri, _ := url.Parse(original)
	uri.Scheme = string(scheme)

	return uri.String()
}

func (m *migu) parseHls(original string, scheme gox.URIScheme) string {
	uri, _ := url.Parse(original)
	uri.Scheme = string(scheme)
	uri.Host = "wshlslive.migucloud.com"

	return uri.String()
}
