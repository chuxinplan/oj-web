package managers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/open-fightcoder/oj-web/common/components"
	"github.com/open-fightcoder/oj-web/data"
	"github.com/open-fightcoder/oj-web/models"
	"github.com/pkg/errors"
)

func AccountLogin(email string, password string) (int64, string, error) {
	account, err := models.AccountGetByEmail(email)
	if err != nil {
		return 0, "", errors.New("登录失败")
	}
	if account == nil {
		return 0, "", errors.New("Email is not exist")
	}
	if account.Password != md5Encode(password) {
		return 0, "", errors.New("Password is wrong")
	}
	user, err := models.GetByAccountId(account.Id)
	if err != nil || user == nil {
		return 0, "", errors.New("登录失败")
	}
	token, err := components.CreateToken(user.Id)
	if err != nil {
		return 0, "", err
	}
	return user.Id, token, nil
}

func AccountRegister(userName string, email string, password string) (int64, error) {
	//TODO 邮箱参数校验,userName校验
	return data.UserRegister(userName, email, md5Encode(password))
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
