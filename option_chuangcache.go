package ula

var _ Option = (*optionChuangcache)(nil)

type optionChuangcache struct {
	domainBase
}

// Chuangcache 配置创世云直播
func Chuangcache(push *domain, pull *domain) *optionChuangcache {
	return &optionChuangcache{
		domainBase: domainBase{
			push: push,
			pull: pull,
		},
	}
}

func (c *optionChuangcache) apply(options *options) {
	options.chuangcache.push = c.push
	options.chuangcache.pull = c.push
	options.ulaType = TypeChuangcache
}
