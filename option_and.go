package ula

var _ Option = (*optionAnd)(nil)

type optionAnd struct {
	endpoint     string
	clientId     string
	clientSecret string
	uid          string
}

// And 配置和直播
func And(clientId string, clientSecret string, uid string) *optionAnd {
	return &optionAnd{
		uid:          uid,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

// AndWithEndpoint 配置和直播
func AndWithEndpoint(endpoint string, clientId string, clientSecret string, uid string) *optionAnd {
	return &optionAnd{
		endpoint:     endpoint,
		uid:          uid,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (a *optionAnd) apply(options *options) {
	if "" != a.endpoint {
		options.migu.endpoint = a.endpoint
	}
	options.and.clientId = a.clientId
	options.and.clientSecret = a.clientSecret
	options.and.uid = a.uid

	options.ulaType = TypeAnd
}
