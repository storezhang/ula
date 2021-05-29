package ula

import (
	"time"
)

var _ Option = (*optionExpired)(nil)

type optionExpired struct {
	expired time.Duration
}

// Expired 配置应用名称
func Expired(expired time.Duration) *optionExpired {
	return &optionExpired{
		expired: expired,
	}
}

func (b *optionExpired) apply(options *options) {
	options.expired = b.expired
}
