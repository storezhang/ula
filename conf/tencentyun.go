package conf

// Tencentyun 腾讯云直播配置
type Tencentyun struct {
	Live

	// BizId 支持400ms左右的超低延迟播放
	BizId int `json:"bizId" yaml:"bizId"`
}
