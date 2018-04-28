package data

import (
	"errors"

	. "github.com/open-fightcoder/oj-web/common/store"
	"github.com/open-fightcoder/oj-web/models"
)

func UserRegister(userName string, email string, password string) (int64, error) {
	session := OrmWeb.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return 0, errors.New("注册失败")
	}
	has, err := session.Where("email = ? ", email).ForUpdate().Get(&models.Account{})
	if err != nil {
		return 0, errors.New("注册失败")
	}
	if has {
		return 0, errors.New("邮箱已存在")
	}
	has, err = session.Where("user_name = ? ", userName).ForUpdate().Get(&models.User{})
	if err != nil {
		return 0, errors.New("注册失败")
	}
	if has {
		return 0, errors.New("用户名已存在")
	}
	account := &models.Account{Email: email, Password: password}
	_, err = session.Insert(account)
	if err != nil {
		session.Rollback()
		return 0, errors.New("注册失败")
	}
	user := &models.User{AccountId: account.Id, UserName: userName}
	_, err = session.Insert(user)
	if err != nil {
		session.Rollback()
		return 0, errors.New("注册失败")
	}
	err = session.Commit()
	if err != nil {
		return 0, errors.New("注册失败")
	}
	return user.Id, nil
}
