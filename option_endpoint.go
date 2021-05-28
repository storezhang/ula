package ula

var _ option = (*optionEndpoint)(nil)

type optionEndpoint struct {
	// 通信端点
	endpoint string
	// 类型
	ulaType Type
}

// Endpoint 配置通信端点
func Endpoint(endpoint string, ulaType Type) *optionEndpoint {
	return &optionEndpoint{
		endpoint: endpoint,
		ulaType:  ulaType,
	}
}

// TencentURL 配置腾讯地址
func TencentURL(url string) *optionEndpoint {
	return Endpoint(url, TypeTencentyun)
}

// AndLiveURL 配置和直播地址
func AndLiveURL(url string) *optionEndpoint {
	return Endpoint(url, TypeAndLive)
}

// ChuangcacheURL 配置创世云地址
func ChuangcacheURL(url string) *optionEndpoint {
	return Endpoint(url, TypeChuangcache)
}

func (b *optionEndpoint) apply(options *options) {
	options.endpoint = b.endpoint
	options.ulaType = b.ulaType
}
