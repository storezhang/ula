package ula

type (
	ulaInternal interface {
		createLive(req *CreateLiveReq, opts *options) (id string, err error)
		getPushUrls(id string, opts *options) (urls []Url, err error)
		getPullCameras(id string, opts *options) (cameras []Camera, err error)
	}
)
