package and

import (
	`fmt`
	`strconv`
	`time`

	`github.com/class100/live-protocol-go`
	`github.com/go-resty/resty/v2`
	`github.com/storezhang/gox`

	`github.com/storezhang/ala/conf`
)

type Live struct {
	config conf.AndLive

	resty *resty.Request
	cache map[string]createAndLiveEventRsp
}

func NewLive(config conf.AndLive, resty *resty.Request) *Live {
	return &Live{
		config: config,

		resty: resty,
		cache: make(map[string]createAndLiveEventRsp),
	}
}

func (l *Live) Create(req *protocol.CreateReq) (id string, err error) {
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
		Name:        req.Title,
		StartTime:   gox.ParseTimestamp(time.Unix(req.StartTime, 0)),
		EndTime:     gox.ParseTimestamp(time.Unix(req.EndTime, 0)),
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

func (l *Live) GetPushUrl(id string) (rsp *protocol.GetPushRsp, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	rsp = new(protocol.GetPushRsp)
	rsp.Index = 1
	rsp.Urls = []*protocol.Url{
		{
			Type: protocol.FormatType_RTMP,
			Link: createRsp.PushUrl,
		},
	}

	return
}

func (l *Live) GetPullUrl(id string) (rsp *protocol.GetPullRsp, err error) {
	var (
		createRsp createAndLiveEventRsp
		ok        bool
	)

	if createRsp, ok = l.cache[id]; !ok {
		return
	}

	rsp = new(protocol.GetPullRsp)
	rsp.Index = 1
	rsp.Cameras = []*protocol.Camera{
		{
			Index: 1,
			Videos: []*protocol.Video{
				{
					Type: protocol.VideoType_ORIGINAL,
					Urls: []*protocol.Url{
						{
							Type: protocol.FormatType_HLS,
							Link: createRsp.Urls[0],
						},
					},
				},
			},
		},
	}

	return
}

func (l *Live) getAndLiveToken() (token string, err error) {
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

func (l *Live) toMap(obj interface{}) (model map[string]string, err error) {
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
