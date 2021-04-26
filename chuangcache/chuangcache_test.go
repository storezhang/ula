package chuangcache

import (
	`fmt`
	`testing`

	`github.com/rs/xid`
	`github.com/storezhang/ala/conf`
)

func TestChuangcache(t *testing.T) {
	chuangcache := NewLive(conf.Chuangcache{
		Live: conf.Live{
			Push: conf.Domain{
				Domain: "push.xxxx.com",
				Key:    "abcd",
			},
			Pull: conf.Domain{
				Domain: "pull.xxxx.com",
				Key:    "efjh",
			},
			Expiration: 6000000,
			Scheme:     "https",
		},
	})

	id := xid.New().String()
	if cameras, err := chuangcache.GetPullCameras(id); nil != err {
		t.FailNow()
	} else {
		fmt.Println(cameras[0])
	}

	if urls, err := chuangcache.GetPushUrls(id); nil != err {
		t.FailNow()
	} else {
		fmt.Println(urls[0])
	}
}
