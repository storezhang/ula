package ula

import (
	`github.com/storezhang/pangu`
)

func init() {
	app := pangu.New()

	if err := app.Provides(New); nil != err {
		panic(err)
	}
}
