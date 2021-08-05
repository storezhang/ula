package ula

const (
	// TypeAnd 和直播
	TypeAnd Type = "and"
	// TypeTencentyun 腾讯云直播
	TypeTencentyun Type = "tencentyun"
	// TypeChuangcache 创世云直播
	TypeChuangcache Type = "chuangcache"
	// TypeMigu 咪咕直播
	TypeMigu Type = "migu"
)

// Type 直播类型
type Type string
