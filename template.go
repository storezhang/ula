package ula

type template struct {
	andLive     executor
	chuangcache executor
	tencentyun  executor
	migu        executor
}

func (t *template) CreateLive(req *CreateLiveReq, opts ...Option) (id string, err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAnd:
		id, err = t.andLive.createLive(req, options)
	case TypeTencentyun:
		id, err = t.tencentyun.createLive(req, options)
	case TypeChuangcache:
		id, err = t.chuangcache.createLive(req, options)
	case TypeMigu:
		id, err = t.migu.createLive(req, options)
	}

	return
}

func (t *template) GetPushUrls(id string, opts ...Option) (urls []Url, err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAnd:
		urls, err = t.andLive.getPushUrls(id, options)
	case TypeTencentyun:
		urls, err = t.tencentyun.getPushUrls(id, options)
	case TypeChuangcache:
		urls, err = t.chuangcache.getPushUrls(id, options)
	case TypeMigu:
		urls, err = t.migu.getPushUrls(id, options)
	}

	return
}

func (t *template) GetPullCameras(id string, opts ...Option) (cameras []Camera, err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAnd:
		cameras, err = t.andLive.getPullCameras(id, options)
	case TypeTencentyun:
		cameras, err = t.tencentyun.getPullCameras(id, options)
	case TypeChuangcache:
		cameras, err = t.chuangcache.getPullCameras(id, options)
	case TypeMigu:
		cameras, err = t.migu.getPullCameras(id, options)
	}

	return
}

func (t *template) Stop(id string, opts ...Option) (success bool, err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAnd:
		success, err = t.andLive.stop(id, options)
	case TypeTencentyun:
		success, err = t.tencentyun.stop(id, options)
	case TypeChuangcache:
		success, err = t.chuangcache.stop(id, options)
	case TypeMigu:
		success, err = t.migu.stop(id, options)
	}

	return
}

func (t *template) GetViewerNum(id string, opts ...Option) (viewerNum int64, err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	switch options.ulaType {
	case TypeAnd:
		viewerNum, err = t.andLive.getViewerNum(id, options)
	case TypeTencentyun:
		viewerNum, err = t.tencentyun.getViewerNum(id, options)
	case TypeChuangcache:
		viewerNum, err = t.chuangcache.getViewerNum(id, options)
	case TypeMigu:
		viewerNum, err = t.migu.getViewerNum(id, options)
	}

	return
}
