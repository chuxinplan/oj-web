package submit

import (
	"errors"

	"time"

	"github.com/open-fightcoder/oj-web/models"
)

func SubmitList(problemId int64, userName string, status int, lang string, currentPage int, perPage int) (map[string]interface{}, error) {
	//TODO 根据userName->userId
	submits, err := models.SubmitGetByConds(problemId, 1, status, lang, currentPage, perPage)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	count, err := models.CountByConds(problemId, 1, status, lang)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	submitMess := map[string]interface{}{
		"list":         submits,
		"current_page": currentPage,
		"total":        count,
	}
	return submitMess, nil
}

func SubmitCommon(problemId int64, userId int64, language string, code string) (map[string]interface{}, error) {
	submit := &models.Submit{ProblemId: problemId, UserId: userId, Language: language, Code: code, SubmitTime: time.Now().Unix()}
	id, err := models.SubmitCreate(submit)
	if err != nil {
		return nil, errors.New("提交失败")
	}
	submitMess := map[string]interface{}{
		"submit_id": id,
		"flag":      1,
	}
	return submitMess, nil
}
func SubmitTest(userId int64, language string, input string, code string) (map[string]interface{}, error) {
	submitTest := &models.SubmitTest{Input: input, UserId: userId, Language: language, Code: code, SubmitTime: time.Now().Unix()}
	id, err := models.SubmitTestCreate(submitTest)
	if err != nil {
		return nil, errors.New("提交失败")
	}
	submitMess := map[string]interface{}{
		"submit_id": id,
		"flag":      2,
	}
	return submitMess, nil
}
func SubmitGetCommon(SubmitId int64) (*models.Submit, error) {
	return models.SubmitGetById(SubmitId)
}
func SubmitGetTest(SubmitId int64) (*models.SubmitTest, error) {
	return models.SubmitTestGetById(SubmitId)
}
