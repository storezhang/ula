package ula

var _ Option = (*optionAndLive)(nil)

type optionAndLive struct {
	// 通信端点
	endpoint string
	// 授权，类似于用户名
	clientId string
	// 授权，类似于密码
	clientSecret string
	// 用户编号，和直播的API有问题，无法登录
	uid int64
}

// AndLive 配置和直播
func AndLive(endpoint string, clientId string, clientSecret string, uid int64) *optionAndLive {
	return &optionAndLive{
		endpoint:     endpoint,
		uid:          uid,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (al *optionAndLive) apply(options *options) {
	options.andLive.clientId = al.clientId
	options.andLive.clientSecret = al.clientSecret
	options.andLive.endpoint = al.endpoint
	options.andLive.uid = al.uid
}
