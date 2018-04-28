package managers

import (
	"crypto/md5"
	"fmt"
	"io"

	"github.com/open-fightcoder/oj-web/common/components"
	"github.com/open-fightcoder/oj-web/models"
	"github.com/pkg/errors"
)

func AccountLogin(email string, password string) (string, error) {
	account, err := models.AccountGetByEmail(email)
	if err != nil {
		return "", errors.New("登录失败")
	}
	if account == nil {
		return "", errors.New("Email is not exist")
	}
	if account.Password != md5Encode(password) {
		return "", errors.New("Password is wrong")
	}
	user, err := models.GetByAccountId(account.Id)
	if err != nil || user == nil {
		return "", errors.New("登录失败")
	}
	token, err := components.CreateToken(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func AccountRegister(userName string, email string, password string) (int64, error) {
	account, err := models.AccountGetByEmail(email)
	if err != nil {
		return 0, fmt.Errorf("get account failure : %s ", err.Error())
	}
	if account != nil {
		return 0, errors.New("Email is exist")
	}
	account = &models.Account{Email: email, Password: md5Encode(password)}
	accountId, err := models.AccountAdd(account)
	if err != nil {
		return 0, errors.New("注册失败")
	}
	insertId, err := models.Create(&models.User{AccountId: accountId, UserName: userName})
	if err != nil {
		return 0, errors.New("注册失败")
	}
	return insertId, nil
}

func md5Encode(password string) string {
	w := md5.New()
	io.WriteString(w, password)
	md5str := string(fmt.Sprintf("%x", w.Sum(nil)))
	return md5str
}
