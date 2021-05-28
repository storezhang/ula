package ula

type ulaTemplate struct {
	andLive     ulaInternal
	chuangcache ulaInternal
	tencentyun  ulaInternal
}

// CreateLive 创建直播信息
func (t *ulaTemplate) CreateLive(req *CreateLiveReq, opts ...option) (id string, err error) {
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

// GetPushUrls 获得推流信息
func (t *ulaTemplate) GetPushUrls(id string, opts ...option) (urls []Url, err error) {
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

// GetPullCameras 获得拉流信息
func (t *ulaTemplate) GetPullCameras(id string, opts ...option) (cameras []Camera, err error) {
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
