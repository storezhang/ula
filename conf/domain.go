package conf

// Domain 域名
type Domain struct {
	// Domain 域名
	Domain string `json:"url" yaml:"url" validate:"required"`
	// Key 鉴权密钥，主要用来生成防盗链地址
	Key string `json:"key" yaml:"key" validate:"required"`
}
