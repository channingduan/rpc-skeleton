package controller

import (
	"github.com/channingduan/rpc/cache"
	"github.com/channingduan/rpc/client"
	"github.com/channingduan/rpc/config"
	"github.com/channingduan/rpc/database"
	"github.com/channingduan/rpc/validator"
)

type Controller struct {
	client    *client.RpcClient
	config    *config.Config
	cache     *cache.Cache
	database  *database.Database
	validator *validator.Validator
}

func Register(config *config.Config) *Controller {

	controller := Controller{
		client:    client.NewClient(config),
		config:    config,
		cache:     cache.Register(&config.CacheConfig),
		database:  database.Register(config),
		validator: validator.NewValidator(),
	}

	return &controller
}
