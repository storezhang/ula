package ula

import (
	"fmt"
	"time"

	"github.com/rs/xid"
	"github.com/storezhang/gox"

	"github.com/storezhang/ula/conf"
)

type Chuangcache struct {
	config conf.Chuangcache

	template ulaTemplate
}

// NewLive 创建创世云直播实现类
func NewChuangcache() (chuangcache *Chuangcache) {
	var config conf.Chuangcache

	chuangcache = &Chuangcache{
		config: config,
	}

	chuangcache.template = ulaTemplate{chuangcache: chuangcache}

	return
}

// CreateLive 创建直播信息
func (l *Chuangcache) CreateLive(req *CreateLiveReq, opts ...option) (id string, err error) {
	return l.template.CreateLive(req, opts...)
}

// GetLivePushFlowInfo 获得推流信息
func (l *Chuangcache) GetLivePushFlowInfo(id string, opts ...option) (urls []Url, err error) {
	return l.template.GetLivePushFlowInfo(id, opts...)
}

// GetLivePullFlowInfo 获得拉流信息
func (l *Chuangcache) GetLivePullFlowInfo(id string, opts ...option) (cameras []Camera, err error) {
	return l.template.GetLivePullFlowInfo(id, opts...)
}

// createLive 创建直播信息
func (l *Chuangcache) createLive(req *CreateLiveReq, opts *options) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

// getLivePushFlowInfo 获得推流信息
func (l *Chuangcache) getLivePushFlowInfo(id string, opts *options) (urls []Url, err error) {
	urls = []Url{
		{
			Type: VideoFormatTypeRtmp,
			Link: l.makeUrl(VideoFormatTypeRtmp, l.config.Push.Domain, id, 1, true),
		},
	}

	return
}

// getLivePullFlowInfo 获得拉流信息
func (l *Chuangcache) getLivePullFlowInfo(id string, opts *options) (cameras []Camera, err error) {
	cameras = []Camera{
		{
			Index: 1,
			Videos: []Video{
				{
					Type: VideoTypeOriginal,
					Urls: []Url{
						{
							Type: VideoFormatTypeRtmp,
							Link: l.makeUrl(VideoFormatTypeRtmp, l.config.Pull.Domain, id, 1, false),
						},
						{
							Type: VideoFormatTypeHls,
							Link: l.makeUrl(VideoFormatTypeHls, l.config.Pull.Domain, id, 1, false),
						},
						{
							Type: VideoFormatTypeFlv,
							Link: l.makeUrl(VideoFormatTypeFlv, l.config.Pull.Domain, id, 1, false),
						},
					},
				},
			}},
	}

	return
}

func (l *Chuangcache) makeUrl(
	formatType VideoFormatType,
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
			l.config.Scheme,
			domain,
			streamName,
			secret,
			expirationTime,
		)
	case VideoFormatTypeHls:
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
