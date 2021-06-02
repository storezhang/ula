package ula

type ulaTemplate struct {
	andLive     ulaInternal
	chuangcache ulaInternal
	tencentyun  ulaInternal
}

func (t *ulaTemplate) CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAndLive:
		id, err = t.andLive.createLive(req, options)
	case TypeTencentyun:
		id, err = t.tencentyun.createLive(req, options)
	}

	return
}

func (t *ulaTemplate) GetPushUrls(id string, opts ...Option) (urls []Url, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAndLive:
		urls, err = t.andLive.getPushUrls(id, options)
	case TypeTencentyun:
		urls, err = t.tencentyun.getPushUrls(id, options)
	case TypeChuangcache:
		urls, err = t.chuangcache.getPushUrls(id, options)
	}

	return
}

func (t *ulaTemplate) GetPullCameras(id string, opts ...Option) (cameras []Camera, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAndLive:
		cameras, err = t.andLive.getPullCameras(id, options)
	case TypeTencentyun:
		cameras, err = t.tencentyun.getPullCameras(id, options)
	case TypeChuangcache:
		cameras, err = t.chuangcache.getPullCameras(id, options)
	}

	return
}

func (t *ulaTemplate) Stop(id string, opts ...Option) (success bool, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAndLive:
		success, err = t.andLive.stop(id, options)
	case TypeTencentyun:
		success, err = t.tencentyun.stop(id, options)
	case TypeChuangcache:
		success, err = t.chuangcache.stop(id, options)
	}

	return
}
