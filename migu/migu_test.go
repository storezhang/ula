package migu_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/storezhang/gox"

	"github.com/storezhang/ula/conf"
	"github.com/storezhang/ula/migu"
	"github.com/storezhang/ula/vo"

	"github.com/storezhang/ula"
)

func initLive() ula.Live {
	return migu.NewLive(conf.Migu{
		Scheme: "http",
		Addr:   "47.111.175.178:9060",
	})
}
func TestMiguLiveCreate(t *testing.T) {
	live := initLive()
	channelID, err := live.Create(vo.Create{
		Title:     fmt.Sprintf("随机title-%s", time.Now().Format("20060102150405")),
		StartTime: gox.Now(),
		EndTime:   gox.ParseTimestamp(time.Now().Add(12 * time.Hour)),
		Extra: map[string]string{
			"subject": fmt.Sprintf("随机subject-%s", time.Now().Format("20060102150405")),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("channel_id", channelID)
}

func TestMiguLiveGetPushUrls(t *testing.T) {
	live := initLive()
	// TODO: channel_id
	channelID := ""
	urls, err := live.GetPushUrls(channelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("urls", urls)
}

func TestMiguLiveGetPullCameras(t *testing.T) {
	live := initLive()
	// TODO: channel_id
	channelID := ""
	urls, err := live.GetPullCameras(channelID)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("urls", urls)
}
