package ula

import (
	`time`
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

func (e *optionExpired) apply(options *options) {
	options.expired = e.expired
}
