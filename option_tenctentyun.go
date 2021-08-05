package ula

var _ Option = (*optionTencentyun)(nil)

type optionTencentyun struct {
	domainBase
}

// Tencentyun 配置腾讯云直播
func Tencentyun(push *domain, pull *domain) *optionTencentyun {
	return &optionTencentyun{
		domainBase: domainBase{
			push: push,
			pull: pull,
		},
	}
}

func (t *optionTencentyun) apply(options *options) {
	options.tencentyun.push = t.push
	options.tencentyun.pull = t.push
	options.ulaType = TypeTencentyun
}
