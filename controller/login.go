package controller

import (
	"context"
	"github.com/channingduan/rpc-skeleton/models"
	"github.com/channingduan/rpc/config"
	"github.com/oscto/ky3k"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Controller) Login(ctx context.Context, request *config.Request, reply *config.Response) error {

	var data LoginRequest
	if err := ky3k.StringToJson(request.Message, &data); err != nil {
		return err
	}
	var user models.User
	c.database.NewDatabase().Where("username = ?", data.Username).First(&user)

	reply.Message = ky3k.JsonToString(user)

	return nil
}
