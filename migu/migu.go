package migu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/storezhang/ula/conf"
	"github.com/storezhang/ula/vo"
)

type live struct {
	config conf.Migu
}

// NewLive 创建咪咕直播实现类
func NewLive(config conf.Migu) *live {
	return &live{
		config: config,
	}
}

// generate token
func (l *live) genToken() string {
	return ""
}

func (l *live) Create(create vo.Create) (id string, err error) {
	if create.Extra == nil {
		return "", ParamMissing
	}
	if _, exist := create.Extra["subject"]; !exist {
		return "", ParamMissSubject
	}

	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/createChannel", l.config.Scheme, l.config.Addr)
	bodyMap := make(map[string]string, len(create.Extra)+3)
	for k, v := range create.Extra {
		bodyMap[k] = v
	}
	bodyMap["title"] = create.Title
	bodyMap["startTime"] = create.StartTime.Time().Format("2006-01-02 15:04:05")
	bodyMap["endTime"] = create.EndTime.Time().Format("2006-01-02 15:04:05")

	bts, _ := json.Marshal(bodyMap)
	body := string(bts)

	var sj *simplejson.Json
	if sj, err = l.httpDo(http.MethodPost, url, bytes.NewBufferString(body)); err != nil {
		return
	}

	id, err = sj.Get("data").Get("channelId").String()

	return
}

func (l *live) GetPushUrls(channelID string) (urls []vo.Url, err error) {
	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/getPushUrl", l.config.Scheme, l.config.Addr)
	body := fmt.Sprintf(`{"channel_id": "%s"}`, channelID)

	var sj *simplejson.Json
	if sj, err = l.httpDo(http.MethodPost, url, bytes.NewBufferString(body)); err != nil {
		return
	}

	var resultBts []byte
	resultBts, err = sj.Get("result").Bytes()
	if err != nil {
		return
	}

	result := new(PushUrl)
	if err = json.Unmarshal(resultBts, result); err != nil {
		return
	}

	urls = make([]vo.Url, len(result.CameraList))
	for i, camera := range result.CameraList {
		urls[i].Link = camera.URL
	}

	return
}

func (l *live) GetPullCameras(channelID string) (cameras []vo.Camera, err error) {
	url := fmt.Sprintf("%s://%s/eduOnlineApi/migu/getPullUrl", l.config.Scheme, l.config.Addr)
	body := fmt.Sprintf(`{"channel_id": "%s"}`, channelID)

	var sj *simplejson.Json
	if sj, err = l.httpDo(http.MethodPost, url, bytes.NewBufferString(body)); err != nil {
		return
	}

	var resultBts []byte
	resultBts, err = sj.Get("result").Bytes()
	if err != nil {
		return
	}

	result := new(PullUrl)
	if err = json.Unmarshal(resultBts, result); err != nil {
		return
	}

	cameras = make([]vo.Camera, len(result.CameraList))
	for i, camera := range result.CameraList {

		videos := make([]vo.Video, len(camera.TranscodeList))
		for _, v := range camera.TranscodeList {
			videos = append(videos, vo.Video{
				Type: transType2VideoType(v.TransType),
				Urls: []vo.Url{
					vo.Url{
						Type: vo.FormatTypeHls,
						Link: v.URLHls,
					},
					vo.Url{
						Type: vo.FormatTypeRtmp,
						Link: v.URLRtmp,
					},
					vo.Url{
						Type: vo.FormatTypeFlv,
						Link: v.URLFlv,
					},
				},
			})
		}

		cameras[i].Videos = videos
		cameras[i].Index = camera.camIndex
	}

	return

}

func (l *live) httpDo(method, url string, body io.Reader) (sj *simplejson.Json, err error) {
	var req *http.Request
	req, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	return l.hasErr(resp.Body)
}

func (l *live) hasErr(reader io.Reader) (sj *simplejson.Json, err error) {
	sj, err = simplejson.NewFromReader(reader)
	if err != nil {
		return
	}

	var ret int
	ret, err = sj.Get("ret").Int()
	if err != nil {
		return
	}

	if ret != 0 {
		var errStr string
		errStr, err = sj.Get("msg").String()
		if err != nil {
			return
		}
		err = errors.New(errStr)
		return
	}

	return
}

func transType2VideoType(transType string) vo.VideoType {
	switch transType {
	case "0":
		return vo.VideoTypeOriginal
	case "1":
		return vo.VideoType360P
	case "2":
		return vo.VideoType480P
	case "3":
		return vo.VideoType720P
	case "4":
		return vo.VideoType1080P
	case "5":
		return vo.VideoTypeAudio
	}
	return vo.VideoType360P
}
