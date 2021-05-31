package ula

var _ Option = (*optionBizId)(nil)

type optionBizId struct {
	bizId int
}

// BizId 配置加速
func BizId(bizId int) *optionBizId {
	return &optionBizId{
		bizId: bizId,
	}
}

func (b *optionBizId) apply(options *options) {
	options.tencentyun.bizId = b.bizId
}
