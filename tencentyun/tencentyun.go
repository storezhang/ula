package tencentyun

import (
	`crypto/md5`
	`fmt`
	`strconv`
	`strings`
	`time`

	`github.com/rs/xid`
	`github.com/storezhang/ala/vo`

	`github.com/storezhang/ala/conf`
)

type live struct {
	config conf.Tencentyun
}

// NewLive 创建腾讯云直播实现类
func NewLive(config conf.Tencentyun) *live {
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
			Link: l.makeUrl(vo.FormatTypeRtmp, l.config.Domain.Push, id, 1, true),
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
							Link: l.makeUrl(vo.FormatTypeRtmp, l.config.Domain.Pull, id, 1, false),
						},
						{
							Type: vo.FormatTypeHls,
							Link: l.makeUrl(vo.FormatTypeHls, l.config.Domain.Pull, id, 1, false),
						},
						{
							Type: vo.FormatTypeFlv,
							Link: l.makeUrl(vo.FormatTypeFlv, l.config.Domain.Pull, id, 1, false),
						},
						{
							Type: vo.FormatTypeRtc,
							Link: l.makeUrl(vo.FormatTypeRtc, l.config.Domain.Pull, id, 1, false),
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
	expirationTime := time.Now().Add(l.config.Expiration).Unix()
	hexExpiration := strings.ToUpper(strconv.FormatInt(expirationTime, 16))
	streamName := fmt.Sprintf("%s-%d", id, camera)

	var key string
	if isPush {
		key = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s", l.config.Key.Push, streamName, hexExpiration))))
	} else {
		key = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%s%s", l.config.Key.Pull, streamName, hexExpiration))))
	}

	switch formatType {
	case vo.FormatTypeRtmp:
		url = fmt.Sprintf(
			"rtmp://%s/live/%s?txSecret=%s&txTime=%s",
			domain,
			streamName,
			key,
			hexExpiration,
		)
	case vo.FormatTypeRtc:
		url = fmt.Sprintf(
			"webrtc://%s/live/%s?txSecret=%s&txTime=%s",
			domain,
			streamName,
			key,
			hexExpiration,
		)
	case vo.FormatTypeFlv:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			hexExpiration,
		)
	case vo.FormatTypeHls:
		url = fmt.Sprintf(
			"%s://%s/live/%s.m3u8?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			hexExpiration,
		)
	default:
		url = fmt.Sprintf(
			"%s://%s/live/%s.flv?txSecret=%s&txTime=%s",
			l.config.Scheme,
			domain,
			streamName,
			key,
			hexExpiration,
		)
	}

	// 超低延时播放：支持400ms左右的超低延迟播放是腾讯云直播播放器的一个特点，它可以用于一些对时延要求极为苛刻的场景，例如远程夹娃娃或者主播连麦等
	// 播放地址需要带防盗链：播放URL不能用普通的CDN URL，必须要带防盗链签名和bizid参数，防盗链签名的计算方法请参见防盗链计算
	// 播放类型需要指定ACC：在调用startPlay函数时，需要指定type为PLAY_TYPE_LIVE_RTMP_ACC，SDK会使用RTMP-UDP协议拉取直播流
	// 该功能有并发播放限制：目前最多同时10路并发播放，避免因为盲目追求低延时而产生不必要的费用损失
	// OBS的延时是不达标的：推流端如果是TXLivePusher，请使用setVideoQuality将quality设置为MAIN_PUBLISHER或者VIDEO_CHAT
	// 该功能按播放时长收费：本功能按照播放时长收费，费用跟拉流的路数有关系，跟音视频流的码率无关，具体价格请参考 价格总览
	if 0 != l.config.Domain.BizId {
		url = fmt.Sprintf("%s&bizid=%d", url, l.config.Domain.BizId)
	}

	return
}
