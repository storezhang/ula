package ula

type domain struct {
	name string
	key  string
}

// NewDomain 创建一个域名
func NewDomain(name string, key string) *domain {
	return &domain{
		name: name,
		key:  key,
	}
}
