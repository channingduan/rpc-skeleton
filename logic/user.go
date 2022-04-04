package logic

import (
	"fmt"
	"github.com/channingduan/rpc/auth"
)

func (l *Logic) UserAuth(id uint) error {

	a := auth.NewAuth(l.config, l.cache)
	token, err := a.CreateToken(id)
	fmt.Println("token: ", token, err)
	if err != nil {
		return err
	}

	return nil
}
