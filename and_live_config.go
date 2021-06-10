package ula

type andLiveConfig struct {
	// 授权，类似于用户名
	clientId string
	// 密码
	clientSecret string
	// 用户编号，和直播的API有问题，无法登录
	uid string
}
