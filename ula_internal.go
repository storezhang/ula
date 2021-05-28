package ula

type (
	ulaInternal interface {
		// createLive 创建直播信息
		createLive(req *CreateLiveReq, opts *options) (id string, err error)
		// getLivePushFlowInfo 获得推流信息
		getLivePushFlowInfo(id string, opts *options) (urls []Url, err error)
		// getLivePullFlowInfo 获得拉流信息
		getLivePullFlowInfo(id string, opts *options) (cameras []Camera, err error)
	}
)
