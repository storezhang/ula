package ula

type domain struct {
	// 地址
	addr string
	// key 鉴权密钥，主要用来生成防盗链地址
	key string
}

// NewDomain 创建一个域名
func NewDomain(addr string, key string) *domain {
	return &domain{
		addr: addr,
		key:  key,
	}
}
