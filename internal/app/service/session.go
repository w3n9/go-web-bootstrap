package service

import (
	"errors"
	"github.com/FromMeToMySelf/go-web-bootstrap/internal/app/model/vo"
)

func Login(param *vo.Login) error {
	if param.Mail == "stringtek@foxmail.com" && param.Password == "123456" {
		return nil
	} else {
		return errors.New("用户名或密码错误")
	}
}
