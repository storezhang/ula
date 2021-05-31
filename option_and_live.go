package ula

var _ option = (*optionAndLive)(nil)

type optionAndLive struct {
	// 通信端点
	endpoint string
	// 授权，类似于用户名
	clientId string
	// 密码
	clientSecret string
	// 用户编号，和直播的API有问题，无法登录
	uid int64
}

// AndLive 配置邮件服务
func AndLive(endpoint string, clientId string, clientSecret string, uid int64) *optionAndLive {
	return &optionAndLive{
		endpoint:     endpoint,
		uid:          uid,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

func (as *optionAndLive) apply(options *options) {
	options.andLive.clientId = as.clientId
	options.andLive.clientSecret = as.clientSecret
	options.andLive.endpoint = as.endpoint
	options.andLive.uid = as.uid
}
