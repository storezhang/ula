package ula

type executor interface {
	createLive(req *CreateLiveReq, options *options) (id string, err error)
	getPushUrls(id string, options *options) (urls []Url, err error)
	getPullCameras(id string, options *options) (cameras []Camera, err error)
	stop(id string, options *options) (success bool, err error)
	getViewerNum(id string, options *options) (viewerNum int64, err error)
}
