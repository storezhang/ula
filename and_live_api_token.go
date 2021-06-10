package ula

type getAndLiveTokenRsp struct {
	andLiveBaseRsp

	// 访问令牌
	AccessToken string `json:"access_token"`
	// 刷新令牌
	RefreshToken string `json:"refresh_token"`
	// 过期时间
	ExpiresIn int64 `json:"expires_in"`
	// 作用域
	Scope string `json:"scope"`
}
