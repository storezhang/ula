package ula

var _ option = (*optionDomain)(nil)

type optionDomain struct {
	// 授权，类似于用户名
	domain string
	// key 鉴权密钥，主要用来生成防盗链地址
	key string
	// 是否是推流
	isPush bool
}

// PushDomain 配置推流域名
func PushDomain(domain string, key string) *optionDomain {
	return &optionDomain{
		domain: domain,
		key:    key,
		isPush: true,
	}
}

// PullDomain 配置推流域名
func PullDomain(domain string, key string) *optionDomain {
	return &optionDomain{
		domain: domain,
		key:    key,
		isPush: false,
	}
}

func (d *optionDomain) apply(options *options) {
	if d.isPush {
		options.pushDomain.domain = d.domain
		options.pushDomain.key = d.key
	} else {
		options.pullDomain.domain = d.domain
		options.pullDomain.key = d.key
	}
}
