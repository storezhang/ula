package chuangcache

import (
	`fmt`
	`time`

	`github.com/rs/xid`
	`github.com/storezhang/gox`
	`github.com/storezhang/ula/vo`

	`github.com/storezhang/ula/conf`
)

type live struct {
	config conf.Chuangcache
}

// NewLive 创建创世云直播实现类
func NewLive(config conf.Chuangcache) *live {
	return &live{
		config: config,
	}
}

func (l *live) Create(_ vo.Create) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

func (l *live) GetPushUrls(id string) (urls []vo.Url, err error) {
	urls = []vo.Url{
		{
			Type: vo.FormatTypeRtmp,
			Link: l.makeUrl(vo.FormatTypeRtmp, l.config.Push.Domain, id, 1, true),
		},
	}

	return
}

func (l *live) GetPullCameras(id string) (cameras []vo.Camera, err error) {
	cameras = []vo.Camera{
		{
			Index: 1,
			Videos: []vo.Video{
				{
					Type: vo.VideoTypeOriginal,
					Urls: []vo.Url{
						{
							Type: vo.FormatTypeRtmp,
							Link: l.makeUrl(vo.FormatTypeRtmp, l.config.Pull.Domain, id, 1, false),
						},
						{
							Type: vo.FormatTypeHls,
							Link: l.makeUrl(vo.FormatTypeHls, l.config.Pull.Domain, id, 1, false),
						},
						{
							Type: vo.FormatTypeFlv,
							Link: l.makeUrl(vo.FormatTypeFlv, l.config.Pull.Domain, id, 1, false),
						},
					},
				},
			}},
	}

	return
}

func (l *live) makeUrl(
	formatType vo.FormatType,
	domain string,
	id string, camera int8,
	isPush bool,
) (url string) {
	expiration := time.Now().Add(l.config.Expiration)
	expirationTime := expiration.Unix()
	expirationString := expiration.Format("2006-01-02T15:04:05Z")
	streamName := fmt.Sprintf("%s-%d", id, camera)

	var (
		token  string
		secret string
	)
	if isPush {
		data := fmt.Sprintf("rtmp://%s/live/%s;%s", l.config.Push.Domain, streamName, expirationString)
		token, _ = gox.Sha256Hmac(data, l.config.Push.Key)
	} else {
		data := fmt.Sprintf("%s/%s/live/%s%d", l.config.Pull.Key, l.config.Pull.Domain, streamName, expirationTime)
		secret, _ = gox.Md5(data)
	}

	switch formatType {
	case vo.FormatTypeRtmp:
		if isPush {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?token=%s?expire=%s",
				domain,
				streamName,
				token,
				expirationString,
			)
		} else {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?secret=%s&timestamp=%s",
				domain,
				streamName,
				secret,
				expirationTime,
			)
		}
	case vo.FormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%s",
			l.config.Scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	case vo.FormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?secret=%s&timestamp=%s",
			l.config.Scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%s",
			l.config.Scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	}

	return
}
