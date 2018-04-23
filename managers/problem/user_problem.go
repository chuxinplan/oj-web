package problem

import (
	"github.com/open-fightcoder/oj-web/models"
	"github.com/pkg/errors"
)

func UserProblemList(user_id int64, current_page int, per_page int) (map[string]interface{}, error) {
	problems, err := models.ProblemGetByUserId(user_id, current_page, per_page)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	count, err := models.ProblemCountByUserId(user_id)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	problemMess := map[string]interface{}{
		"list":         problems,
		"current_page": current_page,
		"total":        count,
	}
	return problemMess, nil
}
