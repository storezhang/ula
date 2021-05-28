package conf

// AndLive 和直播
type AndLive struct {
	// Host Host地址
	Host string `json:"host" yaml:"host" validate:"required,uri"`
	// Id 客户端编号
	Id string `json:"id" yaml:"id" validate:"required"`
	// Secret 客户端密码
	Secret string `json:"secret" yaml:"clientSecret" validate:"required"`
	// Uid 用户编号，因为和直播API的问题，可以得到Uid后硬性配置到系统中，减少复杂度
	Uid int64 `json:"uid" yaml:"uid" validate:"required"`
}
