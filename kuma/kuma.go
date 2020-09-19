package kuma

import (
	"github.com/layer5io/gokit/logger"
	"github.com/mgfeller/common-adapter-library/adapter"
	"github.com/mgfeller/common-adapter-library/config"
)

type KumaAdapter struct {
	adapter.BaseAdapter
}

func New(c config.Handler, l logger.Handler) adapter.Handler {
	return &KumaAdapter{
		adapter.BaseAdapter{Config: c, Log: l},
	}
}
