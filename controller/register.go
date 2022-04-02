package controller

import (
	"context"
	"errors"
	"github.com/channingduan/rpc-skeleton/models"
	"github.com/channingduan/rpc/config"
	"github.com/oscto/ky3k"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Controller) Register(ctx context.Context, request *config.Request, reply *config.Response) error {

	var data RegisterRequest
	if err := c.validator.Bind(request.Message, &data); err != nil {
		return err
	}
	var user models.User
	var total int64
	c.database.NewDatabase().Model(&user).Where("username = ?", data.Username).Count(&total)
	if total > 0 {
		return errors.New("username exists")
	}

	user.Salt = ky3k.RandString(8)
	user.Username = data.Username
	user.Password = ky3k.MarshalMd5(data.Password, user.Salt)
	if err := c.database.NewDatabase().Create(&user).Error; err != nil {
		return err
	}

	return nil
}
