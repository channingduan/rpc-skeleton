package controller

import (
	"context"
	"github.com/channingduan/rpc-skeleton/models"
	"github.com/channingduan/rpc/config"
	"github.com/oscto/ky3k"
)

type ProfileRequest struct {
	UserId uint `json:"user_id"`
}

type ProfileResponse struct {
	UserId   uint   `json:"user_id"`
	Username string `json:"username"`
}

func (c *Controller) Profile(ctx context.Context, request *config.Request, reply *config.Response) error {

	var data ProfileRequest
	if err := c.validator.Bind(request.Message, &data); err != nil {
		return err
	}

	var user models.User
	c.database.NewDatabase().Where("id = ?", data.UserId).Find(&user)

	var result ProfileResponse
	result.UserId = user.ID
	result.Username = user.Username

	reply.Message = ky3k.JsonToString(result)

	return nil
}
