package ula

import (
	`encoding/json`
	`fmt`
	`strconv`
	`strings`
	`sync`
	`time`

	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`
)

var _ executor = (*and)(nil)

type and struct {
	tokenCache  sync.Map
	getCache    sync.Map
	recordCache sync.Map
}

func (a *and) createLive(req *CreateLiveReq, options *options) (id string, err error) {
	var (
		rsp    = new(andApiCreateRsp)
		token  string
		rawRsp *resty.Response
	)

	if token, err = a.getToken(options); nil != err {
		return
	}

	url := fmt.Sprintf("%s/api/v10/events/create.json", options.and.endpoint)
	if rawRsp, err = options.req().SetFormData(map[string]string{
		"client_id":    options.and.clientId,
		"access_token": token,
		"name":         req.Title,
		"starttime":    req.StartTime.Format(),
		"endTime":      req.EndTime.Format(),
		"uid":          options.and.uid,
		"save_video":   "1",
	}).Post(url); nil != err {
		return
	}

	if err = json.Unmarshal(rawRsp.Body(), rsp); nil != err {
		return
	}

	if 0 != rsp.ErrCode {
		err = gox.NewCodeError(gox.ErrorCode(rsp.ErrCode), rsp.ErrMsg, nil)
	} else {
		id = strconv.FormatInt(rsp.Id, 10)
	}

	return
}

func (a *and) getPushUrls(id string, options *options) (urls []Url, err error) {
	var rsp *andApiGetRsp
	if rsp, err = a.get(id, options, true); nil != err {
		return
	}

	urls = []Url{{
		Type: VideoFormatTypeRtmp,
		Link: rsp.PushUrl,
	}}

	return
}

func (a *and) getPullCameras(id string, options *options) (cameras []Camera, err error) {
	var rsp *andApiGetRsp
	if rsp, err = a.get(id, options, true); nil != err {
		return
	}

	if 0 != rsp.ErrCode {
		err = gox.NewCodeError(gox.ErrorCode(rsp.ErrCode), rsp.ErrMsg, nil)
	} else {
		var urls []string
		// 如果直播还没有结束，应该返回拉流地址
		if rsp.EndTime.Time().After(time.Now()) {
			// 取得和直播返回的直播编号，这里做特殊处理，查看返回可以发现规律
			// 20210601210100_7HMMZ6X4
			// http://wshls.live.migucloud.com/live/7HMMZ6X4_C0/playlist.m3u8
			// rtmp://devlivepush.migucloud.com/live/7HMMZ6X4_C0
			urls = []string{fmt.Sprintf("https://wshlslive.migucloud.com/live/%s_C0/playlist.m3u8", rsp.miguId())}
		} else {
			if urls, err = a.recordUrls(id, options); nil != err {
				return
			}
		}
		cameras = []Camera{{
			Index: 1,
			Videos: []Video{{
				Type: VideoTypeOriginal,
				Urls: a.parseLinks(urls),
			}},
		}}
	}

	return
}

func (a *and) stop(id string, options *options) (success bool, err error) {
	var (
		rsp    = new(andApiStopRsp)
		token  string
		rawRsp *resty.Response
	)

	if token, err = a.getToken(options); nil != err {
		return
	}

	url := fmt.Sprintf("%s/api/v10/events/stop.json", options.and.endpoint)
	if rawRsp, err = options.req().SetFormData(map[string]string{
		"client_id":    options.and.clientId,
		"access_token": token,
		"id":           id,
	}).Post(url); nil != err {
		return
	}

	if err = json.Unmarshal(rawRsp.Body(), rsp); nil != err {
		return
	}
	success = 0 == rsp.ErrCode && rsp.Success

	return
}

func (a *and) recordUrls(id string, options *options) (urls []string, err error) {
	var (
		cache interface{}
		ok    bool
		rsp   *andApiGetRsp
	)

	key := id
	if cache, ok = a.recordCache.Load(key); ok {
		urls = cache.([]string)
	} else {
		if rsp, err = a.get(id, options, false); nil != err {
			return
		}
		urls = make([]string, len(rsp.Urls))
		for _, url := range rsp.Urls {
			urls = append(urls, strings.ReplaceAll(url, "http://mgcdn.vod.migucloud.com", "https://mgcdnvod.migucloud.com"))
		}
		a.recordCache.Store(key, urls)
	}

	return
}

func (a *and) get(id string, options *options, useCache bool) (rsp *andApiGetRsp, err error) {
	var (
		cache  interface{}
		ok     bool
		token  string
		rawRsp *resty.Response
	)

	if !useCache {
		a.getCache.Delete(id)
	}

	key := id
	if cache, ok = a.getCache.Load(key); ok {
		rsp = cache.(*andApiGetRsp)
	}
	if nil != rsp {
		return
	}

	if token, err = a.getToken(options); nil != err {
		return
	}

	url := fmt.Sprintf("%s/api/v10/events/get.json", options.and.endpoint)
	if rawRsp, err = options.req().SetQueryParams(map[string]string{
		"client_id":    options.and.clientId,
		"access_token": token,
		"id":           id,
		"uid":          options.and.uid,
	}).Get(url); nil != err {
		return
	}

	rsp = new(andApiGetRsp)
	if err = json.Unmarshal(rawRsp.Body(), rsp); nil != err {
		return
	}
	a.getCache.Store(key, rsp)

	return
}

func (a *and) getToken(options *options) (token string, err error) {
	var (
		cache  interface{}
		ok     bool
		rawRsp *resty.Response
		rsp    = new(andApiTokenRsp)
	)

	key := options.and.clientId
	if cache, ok = a.tokenCache.Load(key); ok {
		var validate bool
		if token, validate = cache.(*andToken).validate(); validate {
			return
		} else {
			a.tokenCache.Delete(key)
		}
	}

	url := fmt.Sprintf("%s/auth/oauth2/access_token", options.and.endpoint)
	if rawRsp, err = options.req().SetFormData(map[string]string{
		"client_id":     options.and.clientId,
		"client_secret": options.and.clientSecret,
		"grant_type":    "client_credentials",
	}).Post(url); nil != err {
		return
	}
	if err = json.Unmarshal(rawRsp.Body(), rsp); nil != err {
		return
	}

	if 1001 == rsp.ErrCode {
		err = &gox.CodeError{Message: rsp.ErrMsg}
	} else {
		token = rsp.AccessToken
		a.tokenCache.Store(key, &andToken{
			accessToken: token,
			expiresIn:   time.Now().Add(time.Duration(1000 * rsp.ExpiresIn)),
		})
	}

	return
}

func (a *and) parseLinks(links []string) (urls []Url) {
	urls = make([]Url, 0, len(links))
	for _, link := range links {
		urls = append(urls, Url{
			Type: VideoFormatTypeHls,
			Link: link,
		})
	}

	return
}

func (a *and) getViewerNum(id string, options *options) (viewerNum int64, err error) {
	return
}
