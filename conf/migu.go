package conf

import (
	`github.com/storezhang/gox`
)

type Migu struct {
	// Scheme 协议
	Scheme gox.URIScheme `default:"http" json:"scheme" yaml:"scheme"`
	Addr   string        `json:"addr" yaml:"addr"`
}
