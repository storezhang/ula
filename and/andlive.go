package and

import (
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/storezhang/gox"
	"github.com/storezhang/ula/vo"

	"github.com/storezhang/ula/conf"
)

type live struct {
	config conf.AndLive

	resty *resty.Request
	cache map[string]createAndLiveEventRsp
}

// NewLive 创建和直播
func NewLive(config conf.AndLive, resty *resty.Request) *live {
	return &live{
		config: config,

		resty: resty,
		cache: make(map[string]createAndLiveEventRsp),
	}
}

func (l *live) Create(create vo.Create) (id string, err error) {
	var (
		andLiveReq map[string]string
		andLiveRsp = new(createAndLiveEventRsp)
		token      string
	)

	if token, err = l.getAndLiveToken(); nil != err {
		return
	}

	params := createAndLiveEventReq{
		ClientId:    l.config.Id,
		AccessToken: token,
		Name:        create.Title,
		StartTime:   create.StartTime,
		EndTime:     create.EndTime,
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

func (l *live) GetPushUrls(id string) (urls []vo.Url, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	urls = []vo.Url{
		{
			Type: vo.FormatTypeRtmp,
			Link: createRsp.PushUrl,
		},
	}

	return
}

func (l *live) GetPullCameras(id string) (cameras []vo.Camera, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	cameras = []vo.Camera{
		{
			Index: "1",
			Videos: []vo.Video{
				{
					Type: vo.VideoTypeOriginal,
					Urls: []vo.Url{
						{
							Type: vo.FormatTypeHls,
							Link: createRsp.Urls[0],
						},
					},
				},
			},
		},
	}

	return
}

func (l *live) getAndLiveToken() (token string, err error) {
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

func (l *live) toMap(obj interface{}) (model map[string]string, err error) {
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
