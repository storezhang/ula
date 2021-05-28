package ula

type ulaTemplate struct {
	andLive        ulaInternal
	chuangcache    ulaInternal
	tencentyunLive ulaInternal
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
		id, err = t.tencentyunLive.createLive(req, options)
	case TypeChuangcache:
		id, err = t.chuangcache.createLive(req, options)
	}

	return
}

// GetLivePushFlowInfo 获得推流信息
func (t *ulaTemplate) GetLivePushFlowInfo(id string, opts ...option) (urls []Url, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	switch options.ulaType {
	case TypeAndLive:
		urls, err = t.andLive.getLivePushFlowInfo(id, options)
	case TypeTencentyun:
		urls, err = t.tencentyunLive.getLivePushFlowInfo(id, options)
	case TypeChuangcache:
		urls, err = t.chuangcache.getLivePushFlowInfo(id, options)
	}

	return
}

// GetLivePullFlowInfo 获得拉流信息
func (t *ulaTemplate) GetLivePullFlowInfo(id string, opts ...option) (cameras []Camera, err error) {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	switch options.ulaType {
	case TypeAndLive:
		cameras, err = t.andLive.getLivePullFlowInfo(id, options)
	case TypeTencentyun:
		cameras, err = t.tencentyunLive.getLivePullFlowInfo(id, options)
	case TypeChuangcache:
		cameras, err = t.chuangcache.getLivePullFlowInfo(id, options)
	}

	return
}
