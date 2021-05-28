package ula

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/rs/xid"
	"github.com/storezhang/gox"

	"github.com/storezhang/ula/conf"
)

type TencentyunLive struct {
	config conf.Tencentyun

	template ulaTemplate
}

// NewLive 创建腾讯云直播实现类
func NewTencentyun() (tencentyunLive *TencentyunLive) {
	var config conf.Tencentyun

	tencentyunLive = &TencentyunLive{
		config: config,
	}

	tencentyunLive.template = ulaTemplate{tencentyunLive: tencentyunLive}

	return
}

// CreateLive 创建直播信息
func (l *TencentyunLive) CreateLive(req *CreateLiveReq, opts ...option) (id string, err error) {
	return l.template.CreateLive(req, opts...)
}

// GetLivePushFlowInfo 获得推流信息
func (l *TencentyunLive) GetLivePushFlowInfo(id string, opts ...option) (urls []Url, err error) {
	return l.template.GetLivePushFlowInfo(id, opts...)
}

// GetLivePullFlowInfo 获得拉流信息
func (l *TencentyunLive) GetLivePullFlowInfo(id string, opts ...option) (cameras []Camera, err error) {
	return l.template.GetLivePullFlowInfo(id, opts...)
}

func (l *TencentyunLive) createLive(req *CreateLiveReq, opts *options) (id string, err error) {
	// 取得和直播返回的直播编号
	id = xid.New().String()

	return
}

func (l *TencentyunLive) getLivePushFlowInfo(id string, opts *options) (urls []Url, err error) {
	urls = []Url{
		{
			Type: VideoFormatTypeRtmp,
			Link: l.makeUrl(VideoFormatTypeRtmp, l.config.Push.Domain, id, 1, true),
		},
	}

	return
}

func (l *TencentyunLive) getLivePullFlowInfo(id string, opts *options) (cameras []Camera, err error) {
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
						{
							Type: VideoFormatTypeRtc,
							Link: l.makeUrl(VideoFormatTypeRtc, l.config.Pull.Domain, id, 1, false),
						},
					},
				},
			}},
	}

	return
}

func (l *TencentyunLive) makeUrl(
	formatType VideoFormatType,
	domain string,
	id string, camera int8,
	isPush bool,
) (url string) {
	expirationTime := time.Now().Add(l.config.Expiration).Unix()
	expirationHex := strings.ToUpper(strconv.FormatInt(expirationTime, 16))
	streamName := fmt.Sprintf("%s-%d", id, camera)

	var key string
	if isPush {
		key, _ = gox.Md5(fmt.Sprintf("%s%s%s", l.config.Push.Key, streamName, expirationHex))
	} else {
		key, _ = gox.Md5(fmt.Sprintf("%s%s%s", l.config.Pull.Key, streamName, expirationHex))
	}

	switch formatType {
	case VideoFormatTypeRtmp:
		url = fmt.Sprintf(
			"rtmp://%s/live/%s?txSecret=%s&txTime=%s",
			domain,
			streamName,
			key,
			expirationHex,
		)
	case VideoFormatTypeRtc:
		url = fmt.Sprintf(
			"webrtc://%s/live/%s?txSecret=%s&txTime=%s",
			domain,
			streamName,
			key,
			expirationHex,
		)
	case VideoFormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			expirationHex,
		)
	case VideoFormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			expirationHex,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			expirationHex,
		)
	}

	// 超低延时播放：支持400ms左右的超低延迟播放是腾讯云直播播放器的一个特点，它可以用于一些对时延要求极为苛刻的场景，例如远程夹娃娃或者主播连麦等
	// 播放地址需要带防盗链：播放URL不能用普通的CDN URL，必须要带防盗链签名和bizid参数，防盗链签名的计算方法请参见防盗链计算
	// 播放类型需要指定ACC：在调用startPlay函数时，需要指定type为PLAY_TYPE_LIVE_RTMP_ACC，SDK会使用RTMP-UDP协议拉取直播流
	// 该功能有并发播放限制：目前最多同时10路并发播放，避免因为盲目追求低延时而产生不必要的费用损失
	// OBS的延时是不达标的：推流端如果是TXLivePusher，请使用setVideoQuality将quality设置为MAIN_PUBLISHER或者VIDEO_CHAT
	// 该功能按播放时长收费：本功能按照播放时长收费，费用跟拉流的路数有关系，跟音视频流的码率无关，具体价格请参考 价格总览
	if 0 != l.config.BizId {
		url = fmt.Sprintf("%s&bizid=%d", url, l.config.BizId)
	}

	return
}
