package ula

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/storezhang/gox"
)

var _ executor = (*chuangcache)(nil)

type chuangcache struct{}

func (c *chuangcache) createLive(_ *CreateLiveReq, _ *options) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

func (c *chuangcache) getPushUrls(id string, options *options) (urls []Url, err error) {
	urls = []Url{{
		Type: VideoFormatTypeRtmp,
		Link: c.makeUrl(VideoFormatTypeRtmp, options.chuangcache.push, true, id, 1, options),
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
				Link: c.makeUrl(VideoFormatTypeRtmp, options.chuangcache.pull, false, id, 1, options),
			}, {
				Type: VideoFormatTypeHls,
				Link: c.makeUrl(VideoFormatTypeHls, options.chuangcache.pull, false, id, 1, options),
			}, {
				Type: VideoFormatTypeFlv,
				Link: c.makeUrl(VideoFormatTypeFlv, options.chuangcache.pull, false, id, 1, options),
			}},
		}},
	}}

	return
}

func (c *chuangcache) stop(_ string, _ *options) (success bool, err error) {
	success = true

	return
}

func (c *chuangcache) getViewerNum(id string, options *options) (viewerNum int, err error) {
	return 0, nil
}

func (c *chuangcache) makeUrl(
	formatType VideoFormatType,
	domain *domain, push bool,
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
	if push {
		data := fmt.Sprintf("rtmp://%s/live/%s;%s", domain.name, streamName, expirationString)
		token, _ = gox.Sha256Hmac(data, domain.key)
	} else {
		data := fmt.Sprintf("%s/%s/live/%s%d", domain.key, domain.name, streamName, expirationTime)
		secret, _ = gox.Md5(data)
	}

	switch formatType {
	case VideoFormatTypeRtmp:
		if push {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?token=%s?expire=%s",
				domain.name,
				streamName,
				token,
				expirationString,
			)
		} else {
			url = fmt.Sprintf(
				"rtmp://%s/live/%s?secret=%s&timestamp=%d",
				domain.name,
				streamName,
				secret,
				expirationTime,
			)
		}
	case VideoFormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%d",
			options.scheme,
			domain.name,
			streamName,
			secret,
			expirationTime,
		)
	case VideoFormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?secret=%s&timestamp=%d",
			options.scheme,
			domain.name,
			streamName,
			secret,
			expirationTime,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?secret=%s&timestamp=%d",
			options.scheme,
			domain.name,
			streamName,
			secret,
			expirationTime,
		)
	}

	return
}
