package ula

var _ option = (*optionSDK)(nil)

type optionSDK struct {
	sdk Type
}

// SDK 配置SDK类型
func SDK(sdk Type) *optionSDK {
	return &optionSDK{
		sdk: sdk,
	}
}

func (s *optionSDK) apply(options *options) {
	options.ulaType = s.sdk
}
