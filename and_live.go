package ula

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/go-resty/resty/v2"
	"github.com/storezhang/gox"

	"github.com/storezhang/ula/conf"
)

type AndLive struct {
	config conf.AndLive

	resty    *resty.Request
	template ulaTemplate

	clientCache sync.Map

	cache map[string]createAndLiveEventRsp
}

// NewAndLive 创建和直播
func NewAndLive(resty *resty.Request) (andLive *AndLive) {
	var conf conf.AndLive

	andLive = &AndLive{
		config: conf,
		resty:  resty,
		cache:  make(map[string]createAndLiveEventRsp),
	}

	andLive.template = ulaTemplate{andLive: andLive}

	return
}

// CreateLive 创建直播信息
func (l *AndLive) CreateLive(req *CreateLiveReq, opts ...option) (id string, err error) {
	return l.template.CreateLive(req, opts...)
}

// GetLivePushFlowInfo 获得推流信息
func (l *AndLive) GetLivePushFlowInfo(id string, opts ...option) (urls []Url, err error) {
	return l.template.GetLivePushFlowInfo(id, opts...)
}

// GetLivePullFlowInfo 获得拉流信息
func (l *AndLive) GetLivePullFlowInfo(id string, opts ...option) (cameras []Camera, err error) {
	return l.template.GetLivePullFlowInfo(id, opts...)
}

// createLive 创建直播信息
func (l *AndLive) createLive(req *CreateLiveReq, opts *options) (id string, err error) {
	var (
		andLiveReq map[string]string
		andLiveRsp = new(createAndLiveEventRsp)
		token      string
	)

	if token, err = l.getAndLiveToken(opts); nil != err {
		return
	}

	params := createAndLiveEventReq{
		ClientId:    l.config.Id,
		AccessToken: token,
		Name:        req.Title,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Uid:         l.config.Uid,
	}
	if andLiveReq, err = l.toMap(params); nil != err {
		return
	}

	url := fmt.Sprintf("%s/api/v10/events/create.json", l.config.Host)
	if _, err = l.resty.SetFormData(andLiveReq).SetResult(andLiveRsp).Post(url); nil != err {
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

// getLivePushFlowInfo 获得推流信息
func (l *AndLive) getLivePushFlowInfo(id string, opts *options) (urls []Url, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	urls = []Url{
		{
			Type: VideoFormatTypeRtmp,
			Link: createRsp.PushUrl,
		},
	}

	return
}

// getLivePullFlowInfo 获得拉流信息
func (l *AndLive) getLivePullFlowInfo(id string, opts *options) (cameras []Camera, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	cameras = []Camera{
		{
			Index: 1,
			Videos: []Video{
				{
					Type: VideoTypeOriginal,
					Urls: []Url{
						{
							Type: VideoFormatTypeHls,
							Link: createRsp.Urls[0],
						},
					},
				},
			},
		},
	}

	return
}

func (l *AndLive) getAndLiveToken(opts *options) (token string, err error) {
	var (
		andLiveReq map[string]string
		rsp        = new(getAndLiveTokenRsp)
	)

	params := &getAndLiveTokenReq{
		ClientId:     l.config.Id,
		ClientSecret: l.config.Secret,
		GrantType:    "client_credentials",
	}
	if andLiveReq, err = l.toMap(params); nil != err {
		return
	}

	url := fmt.Sprintf("%v/auth/oauth2/access_token", l.config.Host)
	if _, err = l.resty.SetFormData(andLiveReq).SetResult(rsp).Post(url); nil != err {
		return
	}

	if 1001 == rsp.ErrCode {
		err = &gox.CodeError{Message: rsp.ErrMsg}

		return
	}

	token = rsp.AccessToken

	return
}

func (l *AndLive) toMap(obj interface{}) (model map[string]string, err error) {
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
