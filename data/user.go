package data

import (
	"errors"

	. "github.com/open-fightcoder/oj-web/common/store"
	"github.com/open-fightcoder/oj-web/models"
)

func UserRegister(userName string, email string, password string) error {
	session := OrmWeb.NewSession()
	defer session.Close()

	err := session.Begin()
	has, err := session.Where("email = ? for update", email).Get(&models.Account{})
	if err != nil {
		return errors.New("注册失败")
	}
	if has {
		return errors.New("邮箱已存在")
	}
	has, err = session.Where("user_name = ? for update", userName).Get(&models.User{})
	if err != nil {
		return errors.New("注册失败")
	}
	if has {
		return errors.New("用户名已存在")
	}
	account := &models.Account{Email: email, Password: password}
	_, err = session.Insert(account)
	if err != nil {
		session.Rollback()
		return errors.New("注册失败")
	}
	user := &models.User{AccountId: account.Id, UserName: userName}
	_, err = session.Insert(user)
	if err != nil {
		session.Rollback()
		return errors.New("注册失败")
	}
	err = session.Commit()
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}
