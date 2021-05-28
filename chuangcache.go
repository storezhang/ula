package ula

import (
	`fmt`
	`time`

	`github.com/rs/xid`
	`github.com/storezhang/gox`
)

type Chuangcache struct {
	template ulaTemplate
}

// NewChuangcache 创建创世云直播实现类
func NewChuangcache() (chuangcache *Chuangcache) {
	chuangcache = &Chuangcache{}
	chuangcache.template = ulaTemplate{chuangcache: chuangcache}

	return
}

func (c *Chuangcache) CreateLive(req *CreateLiveReq, opts ...option) (id string, err error) {
	return c.template.CreateLive(req, opts...)
}

func (c *Chuangcache) GetPushUrls(id string, opts ...option) (urls []Url, err error) {
	return c.template.GetPushUrls(id, opts...)
}

func (c *Chuangcache) GetPullCameras(id string, opts ...option) (cameras []Camera, err error) {
	return c.template.GetPullCameras(id, opts...)
}

// createLive 创建直播信息
func (c *Chuangcache) createLive(_ *CreateLiveReq, _ *options) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

func (c *Chuangcache) getPushUrls(id string, options *options) (urls []Url, err error) {
	urls = []Url{{
		Type: VideoFormatTypeRtmp,
		Link: c.makeUrl(VideoFormatTypeRtmp, options.pushDomain.domain, id, 1, true, options),
	}}

	return
}

func (c *Chuangcache) getPullCameras(id string, options *options) (cameras []Camera, err error) {
	cameras = []Camera{{
		Index: 1,
		Videos: []Video{{
			Type: VideoTypeOriginal,
			Urls: []Url{{
				Type: VideoFormatTypeRtmp,
				Link: c.makeUrl(VideoFormatTypeRtmp, options.pullDomain.domain, id, 1, false, options),
			}, {
				Type: VideoFormatTypeHls,
				Link: c.makeUrl(VideoFormatTypeHls, options.pullDomain.domain, id, 1, false, options),
			}, {
				Type: VideoFormatTypeFlv,
				Link: c.makeUrl(VideoFormatTypeFlv, options.pullDomain.domain, id, 1, false, options),
			}},
		}},
	}}

	return
}

func (c *Chuangcache) makeUrl(
	formatType VideoFormatType,
	domain string,
	id string, camera int8,
	isPush bool,
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
	if isPush {
		data := fmt.Sprintf("rtmp://%s/live/%s;%s", options.pushDomain.domain, streamName, expirationString)
		token, _ = gox.Sha256Hmac(data, options.pushDomain.key)
	} else {
		data := fmt.Sprintf("%s/%s/live/%s%d", options.pullDomain.key, options.pullDomain.domain, streamName, expirationTime)
		secret, _ = gox.Md5(data)
	}

	switch formatType {
	case VideoFormatTypeRtmp:
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
	case VideoFormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%s",
			options.scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	case VideoFormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?secret=%s&timestamp=%s",
			options.scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%s",
			options.scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	}

	return
}
