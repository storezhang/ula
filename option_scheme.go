package ula

import (
	`github.com/storezhang/gox`
)

var _ Option = (*optionScheme)(nil)

type optionScheme struct {
	scheme gox.URIScheme
}

// Scheme 配置协议
func Scheme(scheme gox.URIScheme) *optionScheme {
	return &optionScheme{
		scheme: scheme,
	}
}

func (s *optionScheme) apply(options *options) {
	options.scheme = s.scheme
}
