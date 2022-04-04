package logic

import (
	"github.com/channingduan/rpc/cache"
	"github.com/channingduan/rpc/config"
)

type Logic struct {
	config *config.Config
	cache  *cache.Cache
}

func NewLogic(config *config.Config, cache *cache.Cache) *Logic {
	return &Logic{
		config: config,
		cache:  cache,
	}
}
