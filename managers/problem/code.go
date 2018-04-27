package problem

import (
	"errors"

	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/models"
)

func CodeGet(userId int64, problemId int64) (map[string]interface{}, error) {
	userCode, err := models.UserCodeGetUserCode(userId, problemId)
	if err != nil {
		return nil, errors.New("Get code error!")
	}
	codeMess := map[string]interface{}{
		"code":     managers.GetCode(userCode.SaveCode),
		"language": userCode.Language,
	}
	return codeMess, nil
}

func CodeSet(problemId int64, userId int64, saveCode string, language string) error {
	code, err := models.UserCodeGetUserCode(userId, problemId)
	if err != nil {
		return err
	}
	var errorRet error
	codePath, err := managers.SaveCode(saveCode)
	if err != nil {
		return err
	}
	if code == nil {
		userCode := &models.UserCode{ProblemId: problemId, UserId: userId, SaveCode: codePath, Language: language}
		_, errorRet = models.UserCodeCreate(userCode)
	} else {
		code.SaveCode = codePath
		code.Language = language
		errorRet = models.UserCodeUpdate(code)
	}
	return errorRet
}
