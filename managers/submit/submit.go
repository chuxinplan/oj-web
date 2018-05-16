package submit

import (
	"errors"

	"time"

	"github.com/open-fightcoder/oj-web/common/components"
	"github.com/open-fightcoder/oj-web/common/g"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/models"
)

func SubmitList(problemId int64, userName string, status int, lang string, currentPage int, perPage int) (map[string]interface{}, error) {
	submits, err := models.SubmitGetByConds(problemId, 0, status, lang, currentPage, perPage)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	count, err := models.SubmitCountByConds(problemId, 0, status, lang)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var submitLists []map[string]interface{}
	for _, v := range submits {
		submitTime := time.Unix(v.SubmitTime, 0).Format("2006-01-02 15:04:05")
		user, err := models.GetById(v.UserId)
		if err != nil {
			return nil, errors.New("查询失败")
		}
		problem, err := models.ProblemGetById(v.ProblemId)
		if err != nil {
			return nil, errors.New("查询失败")
		}
		projects := make(map[string]interface{})
		projects["problem_id"] = problem.Id
		projects["problem_name"] = problem.Title
		projects["user_id"] = user.Id
		projects["user_name"] = user.UserName
		projects["nick_name"] = user.NickName
		projects["status"] = v.Result
		projects["memory_cost"] = v.RunningMemory
		projects["time_cost"] = v.RunningTime
		projects["lang"] = v.Language
		projects["submit_id"] = v.Id
		projects["time"] = submitTime
		submitLists = append(submitLists, projects)
	}
	submitMess := map[string]interface{}{
		"list":         submitLists,
		"current_page": currentPage,
		"total":        count,
	}
	return submitMess, nil
}

func IsInOj(userId int64) bool {
	flag := false
	ojIds := g.Conf().Problem.UserId
	for _, v := range ojIds {
		if v == userId {
			flag = true
			break
		}
	}
	return flag
}

func SubmitCommon(problemId int64, userId int64, language string, code string) (map[string]interface{}, error) {
	codePath, err := managers.SaveSubmitCode(code, userId, language)
	if err != nil {
		return nil, err
	}
	submit := &models.Submit{ProblemId: problemId, Result: 1, UserId: userId, Language: language, Code: codePath, SubmitTime: time.Now().Unix()}
	id, err := models.SubmitCreate(submit)
	if err != nil {
		return nil, errors.New("提交失败")
	}
	problem, err := models.ProblemGetById(problemId)
	if err != nil {
		return nil, errors.New("提交失败")
	}
	sendMess := &components.SendMess{"default", id}
	var flag bool
	if IsInOj(problem.UserId) {
		flag = components.Send("vjudger", sendMess)
	} else {
		flag = components.Send("judge", sendMess)
	}
	if flag == false {
		return nil, errors.New("提交失败")
	}
	submitMess := map[string]interface{}{
		"submit_id": id,
	}
	return submitMess, nil
}
func SubmitTest(userId int64, language string, input string, code string) (map[string]interface{}, error) {
	codePath, err := managers.SaveSubmitCode(code, userId, language)
	if err != nil {
		return nil, err
	}
	submitTest := &models.SubmitTest{Input: input, Result: 1, UserId: userId, Language: language, Code: codePath, SubmitTime: time.Now().Unix()}
	id, err := models.SubmitTestCreate(submitTest)
	if err != nil {
		return nil, errors.New("提交失败")
	}
	sendMess := &components.SendMess{"test", id}
	flag := components.Send("judge", sendMess)
	if flag == false {
		return nil, errors.New("提交失败")
	}
	submitMess := map[string]interface{}{
		"submit_id": id,
	}
	return submitMess, nil
}
func SubmitGetCommon(submitId int64, currentId int64) (map[string]interface{}, error) {
	submit, err := models.SubmitGetById(submitId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	problem, err := models.ProblemGetById(submit.ProblemId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	code := "你无权查看"
	if currentId == submit.UserId {
		code, err = managers.GetCode(submit.Code)
		if err != nil {
			return nil, err
		}
	}
	submitMess := map[string]interface{}{
		"problem_id":   problem.Id,
		"problem_name": problem.Title,
		"status":       submit.Result,
		"memory_cost":  submit.RunningMemory,
		"time_cost":    submit.RunningTime,
		"lang":         submit.Language,
		"code":         code,
		"time":         time.Unix(submit.SubmitTime, 0).Format("2006-01-02 15:04:05"),
		"result_des":   submit.ResultDes,
	}
	return submitMess, nil
}
func SubmitGetTest(SubmitId int64) (map[string]interface{}, error) {
	submit, err := models.SubmitTestGetById(SubmitId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	submitMess := map[string]interface{}{
		"status":      submit.Result,
		"memory_cost": submit.RunningMemory,
		"time_cost":   submit.RunningTime,
		"lang":        submit.Language,
		"result_des":  submit.ResultDes,
		"time":        time.Unix(submit.SubmitTime, 0).Format("2006-01-02 15:04:05"),
	}
	return submitMess, nil
}
