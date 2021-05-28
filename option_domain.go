package ula

import (
	"github.com/storezhang/gox"
)

var _ option = (*optionDomain)(nil)

type optionDomain struct {
	// 授权，类似于用户名
	domain string
	// key 鉴权密钥，主要用来生成防盗链地址
	key string
}
