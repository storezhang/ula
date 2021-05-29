package ula

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/storezhang/gox"
)

// NewResty Resty客户端
type andLive struct {
	resty    *resty.Request
	template ulaTemplate
}

// NewAndLive 创建和直播
func NewAndLive(resty *resty.Request) (live *andLive) {
	live = &andLive{
		resty: resty,
	}

	// live.resty.SetProxy("socks5://192.168.178.178:1080")
	live.template = ulaTemplate{andLive: live}

	return
}

func (a *andLive) CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error) {
	return a.template.CreateLive(req, opts...)
}

func (a *andLive) GetPushUrls(id string, opts ...Option) (urls []Url, err error) {
	return a.template.GetPushUrls(id, opts...)
}

func (a *andLive) GetPullCameras(id string, opts ...Option) (cameras []Camera, err error) {
	return a.template.GetPullCameras(id, opts...)
}

func (a *andLive) createLive(req *CreateLiveReq, options *options) (id string, err error) {
	var (
		andLiveReq map[string]string
		andLiveRsp = new(createAndLiveEventRsp)
		token      string
	)

	if token, err = a.getAndLiveToken(options); nil != err {
		return
	}

	params := createAndLiveEventReq{
		ClientId:    options.andLive.clientId,
		AccessToken: token,
		Name:        req.Title,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Uid:         options.andLive.uid,
	}
	if andLiveReq, err = a.toMap(params); nil != err {
		return
	}

	url := fmt.Sprintf("%s/api/v10/events/create.json", options.andLive.endpoint)
	if _, err = a.resty.SetFormData(andLiveReq).SetResult(andLiveRsp).Post(url); nil != err {
		return
	}

	if 0 != andLiveRsp.ErrCode {
		err = gox.NewCodeError(gox.ErrorCode(andLiveRsp.ErrCode), andLiveRsp.ErrMsg, nil)

		return
	}

	// 取得和直播返回的直播编号
	id = strconv.FormatInt(andLiveRsp.Id, 10)

	return
}

func (a *andLive) getPushUrls(liveId string, options *options) (urls []Url, err error) {
	var andLiveRsp *getAndLiveEventRsp

	if andLiveRsp, err = a.getAndLiveEvent(liveId, options); nil != err {
		return
	}

	if nil != andLiveRsp {
		urls = []Url{
			{
				Type: VideoFormatTypeRtmp,
				Link: andLiveRsp.PushUrl,
			},
		}
	}

	return
}

// getLivePullFlowInfo 获得拉流信息
func (a *andLive) getPullCameras(liveId string, options *options) (cameras []Camera, err error) {
	var andLiveRsp *getAndLiveEventRsp

	if andLiveRsp, err = a.getAndLiveEvent(liveId, options); nil != err {
		return
	}

	if nil != andLiveRsp && 0 != len(andLiveRsp.Urls) {
		cameras = []Camera{
			{
				Index: 1,
				Videos: []Video{
					{
						Type: VideoTypeOriginal,
						Urls: []Url{
							{
								Type: VideoFormatTypeHls,
								Link: andLiveRsp.Urls[0],
							},
						},
					},
				},
			},
		}
	}

	return
}

func (a *andLive) getAndLiveEvent(liveId string, options *options) (andLiveRsp *getAndLiveEventRsp, err error) {
	var (
		andLiveReq map[string]string
		id         int64
		token      string
	)

	if id, err = strconv.ParseInt(liveId, 10, 64); nil != err {
		return
	}

	params := &getAndLiveEventReq{
		ClientId:    options.andLive.clientId,
		AccessToken: token,
		Id:          id,
	}

	if andLiveReq, err = a.toMap(params); nil != err {
		return
	}

	andLiveRsp = new(getAndLiveEventRsp)
	url := fmt.Sprintf("%v/api/v10/events/get.json", options.andLive.endpoint)
	if _, err = a.resty.SetQueryParams(andLiveReq).SetResult(andLiveRsp).Get(url); nil != err {
		return
	}

	if 0 != andLiveRsp.ErrCode {
		err = gox.NewCodeError(gox.ErrorCode(andLiveRsp.ErrCode), andLiveRsp.ErrMsg, nil)

		return
	}

	return
}

func (a *andLive) getAndLiveToken(options *options) (token string, err error) {
	var (
		andLiveReq map[string]string
		rsp        = new(getAndLiveTokenRsp)
	)

	params := &getAndLiveTokenReq{
		ClientId:     options.andLive.clientId,
		ClientSecret: options.andLive.clientSecret,
		GrantType:    "client_credentials",
	}

	if andLiveReq, err = a.toMap(params); nil != err {
		return
	}

	url := fmt.Sprintf("%v/auth/oauth2/access_token", options.andLive.endpoint)
	if _, err = a.resty.SetFormData(andLiveReq).SetResult(rsp).Post(url); nil != err {
		return
	}

	if 1001 == rsp.ErrCode {
		err = &gox.CodeError{Message: rsp.ErrMsg}

		return
	}

	token = rsp.AccessToken

	return
}

func (a *andLive) toMap(obj interface{}) (model map[string]string, err error) {
	var flattenParams map[string]interface{}

	model = make(map[string]string)
	if flattenParams, err = gox.StructToMap(obj); nil != err {
		return
	}
	if flattenParams, err = gox.Flatten(flattenParams, "", gox.DotStyle); nil != err {
		return
	}

	for key, value := range flattenParams {
		model[key] = fmt.Sprintf("%v", value)
	}

	return
}
