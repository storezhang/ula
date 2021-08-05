package ula

var _ Option = (*optionMigu)(nil)

type optionMigu struct {
	endpoint  string
	uid       string
	secretId  string
	secretKey string
}

// Migu 配置咪咕直播
func Migu(uid string, secretId string, secretKey string) *optionMigu {
	return &optionMigu{
		uid:       uid,
		secretId:  secretId,
		secretKey: secretKey,
	}
}

// MiguWithEndpoint 配置咪咕直播
func MiguWithEndpoint(endpoint string, uid string, secretId string, secretKey string) *optionMigu {
	return &optionMigu{
		endpoint:  endpoint,
		uid:       uid,
		secretId:  secretId,
		secretKey: secretKey,
	}
}

func (m *optionMigu) apply(options *options) {
	if "" != m.endpoint {
		options.migu.endpoint = m.endpoint
	}
	options.migu.uid = m.uid
	options.migu.secretId = m.secretId
	options.migu.secretKey = m.secretKey

	options.ulaType = TypeMigu
}
