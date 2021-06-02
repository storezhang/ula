package ula

import (
	`fmt`
	`time`

	`github.com/rs/xid`
	`github.com/storezhang/gox`
)

var (
	_ Ula         = (*chuangcache)(nil)
	_ ulaInternal = (*chuangcache)(nil)
)

type chuangcache struct {
	template ulaTemplate
}

// NewChuangcache 创建创世云直播实现类
func NewChuangcache() (live *chuangcache) {
	live = &chuangcache{}
	live.template = ulaTemplate{chuangcache: live}

	return
}

func (c *chuangcache) CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error) {
	return c.template.CreateLive(req, opts...)
}

func (c *chuangcache) GetPushUrls(id string, opts ...Option) (urls []Url, err error) {
	return c.template.GetPushUrls(id, opts...)
}

func (c *chuangcache) GetPullCameras(id string, opts ...Option) (cameras []Camera, err error) {
	return c.template.GetPullCameras(id, opts...)
}

func (c *chuangcache) Stop(id string, opts ...Option) (success bool, err error) {
	return c.template.Stop(id, opts...)
}

func (c *chuangcache) createLive(_ *CreateLiveReq, _ *options) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

func (c *chuangcache) getPushUrls(id string, options *options) (urls []Url, err error) {
	urls = []Url{{
		Type: VideoFormatTypeRtmp,
		Link: c.makeUrl(VideoFormatTypeRtmp, options.pushDomain, id, 1, options),
	}}

	return
}

func (c *chuangcache) getPullCameras(id string, options *options) (cameras []Camera, err error) {
	cameras = []Camera{{
		Index: 1,
		Videos: []Video{{
			Type: VideoTypeOriginal,
			Urls: []Url{{
				Type: VideoFormatTypeRtmp,
				Link: c.makeUrl(VideoFormatTypeRtmp, options.pullDomain, id, 1, options),
			}, {
				Type: VideoFormatTypeHls,
				Link: c.makeUrl(VideoFormatTypeHls, options.pullDomain, id, 1, options),
			}, {
				Type: VideoFormatTypeFlv,
				Link: c.makeUrl(VideoFormatTypeFlv, options.pullDomain, id, 1, options),
			}},
		}},
	}}

	return
}

func (c *chuangcache) stop(_ string, _ *options) (success bool, err error) {
	success = true

	return
}

func (c *chuangcache) makeUrl(
	formatType VideoFormatType,
	domain optionDomain,
	id string, camera int8,
	options *options,
) (url string) {
	expiration := time.Now().Add(options.expired)
	expirationTime := expiration.Unix()
	expirationString := expiration.Format("2006-01-02T15:04:05Z")
	streamName := fmt.Sprintf("%s-%d", id, camera)

	var (
		token  string
		secret string
	)
	if domain.isPush {
		data := fmt.Sprintf("rtmp://%s/live/%s;%s", domain.domain, streamName, expirationString)
		token, _ = gox.Sha256Hmac(data, domain.key)
	} else {
		data := fmt.Sprintf("%s/%s/live/%s%d", domain.key, domain.domain, streamName, expirationTime)
		secret, _ = gox.Md5(data)
	}

	switch formatType {
	case VideoFormatTypeRtmp:
		if domain.isPush {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?token=%s?expire=%s",
				domain.domain,
				streamName,
				token,
				expirationString,
			)
		} else {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?secret=%s&timestamp=%d",
				domain.domain,
				streamName,
				secret,
				expirationTime,
			)
		}
	case VideoFormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%d",
			options.scheme,
			domain.domain,
			streamName,
			secret,
			expirationTime,
		)
	case VideoFormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?secret=%s&timestamp=%d",
			options.scheme,
			domain.domain,
			streamName,
			secret,
			expirationTime,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%d",
			options.scheme,
			domain.domain,
			streamName,
			secret,
			expirationTime,
		)
	}

	return
}
