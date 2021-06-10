package ula

var _ Option = (*optionMigu)(nil)

type optionMigu struct {
	token string
	key   string
	iv    string
}

// Migu 配置咪咕直播
func Migu(token string, key string, iv string) *optionMigu {
	return &optionMigu{
		token: token,
		key:   key,
		iv:    iv,
	}
}

func (al *optionMigu) apply(options *options) {
	options.migu.token = al.token
	options.migu.key = al.key
	options.migu.iv = al.iv
}
