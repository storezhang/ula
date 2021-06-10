package ula

type miguConfig struct {
	// 授权验证
	token string `validate:"required"`
	// 加密
	iv string `validate:"required"`
	// 加密密钥
	key string `validate:"required"`
}
