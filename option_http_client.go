package ula

import (
	`github.com/go-resty/resty/v2`
)

var _ Option = (*optionHttpClient)(nil)

type optionHttpClient struct {
	client *resty.Client
}

// HttpClient 配置Http客户端
func HttpClient(client *resty.Client) *optionHttpClient {
	return &optionHttpClient{
		client: client,
	}
}

func (hc *optionHttpClient) apply(options *options) {
	options.resty = hc.client
}
