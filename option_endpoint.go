package ula

import (
	`fmt`

	`github.com/storezhang/gox`
)

var _ Option = (*optionEndpoint)(nil)

type optionEndpoint struct {
	endpoint string
}

// Endpoint 配置通信端点
func Endpoint(scheme gox.URIScheme, host string, port int) *optionEndpoint {
	return &optionEndpoint{
		endpoint: fmt.Sprintf("%s://%s:%d", scheme, host, port),
	}
}

// EndpointHttp 配置Http端点
func EndpointHttp(host string) *optionEndpoint {
	return Endpoint(gox.URISchemeHttp, host, 80)
}

// EndpointHttps 配置Https端点
func EndpointHttps(host string) *optionEndpoint {
	return Endpoint(gox.URISchemeHttps, host, 443)
}

// EndpointURI 配置通信端点
func EndpointURI(uri string) *optionEndpoint {
	return &optionEndpoint{
		endpoint: uri,
	}
}

func (d *optionEndpoint) apply(options *options) {
	options.endpoint = d.endpoint
}
