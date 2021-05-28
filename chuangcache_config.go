package ula

import (
	"fmt"
)

type chuangcacheConfig struct {
	// 授权，相当于用户名
	ak string `validate:"required"`
	// 授权，相当于密码
	sk string `validate:"required"`
	// 应用
	appKey string `validate:"required"`
}

func (cc *chuangcacheConfig) key() string {
	return fmt.Sprintf("%s-%s", cc.ak, cc.sk)
}
