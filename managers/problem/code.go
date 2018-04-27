package problem

import (
	"errors"

	"github.com/open-fightcoder/oj-web/models"
)

func CodeGet(userId int64, problemId int64) (string, error) {
	code := ""
	userCode, err := models.UserCodeGetUserCode(userId, problemId)
	if err != nil {
		return "", errors.New("Get code error!")
	}
	if userCode != nil {
		code = userCode.SaveCode
	}
	return code, nil
}

func CodeSet(problemId int64, userId int64, saveCode string, language string) error {
	userCode := &models.UserCode{ProblemId: problemId, UserId: userId, SaveCode: saveCode, Language: language}
	_, err := models.UserCodeCreate(userCode)
	return err
}
