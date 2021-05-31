package ula

type (
	getAndLiveTokenReq struct {
		// 客户端编号
		ClientId string `json:"client_id" validate:"required"`
		// 客户端密钥
		ClientSecret string `json:"client_secret" validate:"required"`
		// 固定值，client_credentials
		GrantType string `json:"grant_type,string" validate:"required"`
	}

	getAndLiveTokenRsp struct {
		baseAndLiveRsp

		// 访问令牌
		AccessToken string `json:"access_token"`
		// 刷新令牌
		RefreshToken string `json:"refresh_token"`
		// 过期时间
		ExpiresIn int64 `json:"expires_in"`
		// 作用域
		Scope string `json:"scope"`
	}
)
