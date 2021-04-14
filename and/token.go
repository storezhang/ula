package and

type (
	getAndLiveTokenReq struct {
		// ClientId 客户端 id
		ClientId string `json:"client_id" validate:"required"`
		// ClientSecret 客户端密钥
		ClientSecret string `json:"client_secret" validate:"required"`
		// GrantType 固定值 client_credentials
		GrantType string `json:"grant_type,string" validate:"required"`
	}

	getAndLiveTokenRsp struct {
		baseAndLiveRsp

		// AccessToken 访问令牌
		AccessToken string `json:"access_token"`
		// RefreshToken 刷新令牌
		RefreshToken string `json:"refresh_token"`
		// ExpiresIn 过期时间
		ExpiresIn int64 `json:"expires_in"`
		// Scope 作用域
		Scope string `json:"scope"`
	}
)
