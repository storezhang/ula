package ula

type andLiveConfig struct {
	// 通信端点
	endpoint string
	// 授权，类似于用户名
	clientId string
	// 用户编号，和直播的API有问题，无法登录
	uid int64
}
