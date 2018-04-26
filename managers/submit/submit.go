package submit

import (
	"errors"

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

func SubmitCommon() {

}
