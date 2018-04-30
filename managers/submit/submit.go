package submit

import (
	"errors"

	"time"

	"github.com/open-fightcoder/oj-web/common/components"
	"github.com/open-fightcoder/oj-web/managers"
	"github.com/open-fightcoder/oj-web/models"
)

func SubmitList(problemId int64, userName string, status int, lang string, currentPage int, perPage int) (map[string]interface{}, error) {
	//TODO 根据userName->userId
	submits, err := models.SubmitGetByConds(problemId, 0, status, lang, currentPage, perPage)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	count, err := models.SubmitCountByConds(problemId, 0, status, lang)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var submitLists []map[string]interface{}
	for i := 0; i < len(submits); i++ {
		submitTime := time.Unix(submits[i].SubmitTime, 0).Format("2006-01-02 15:04:05")
		projects := make(map[string]interface{})
		//TODO
		projects["problem_name"] = "测试"
		projects["user_id"] = 1
		projects["user_name"] = "哈哈"
		projects["status"] = submits[i].Result
		projects["memory_cost"] = submits[i].RunningMemory
		projects["time_cost"] = submits[i].RunningTime
		projects["lang"] = submits[i].Language
		projects["submit_id"] = submits[i].Id
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
	sendMess := &components.SendMess{"default", id}
	flag := components.Send("judge", sendMess)
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
func SubmitGetCommon(SubmitId int64) (map[string]interface{}, error) {
	submit, err := models.SubmitGetById(SubmitId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	code, err := managers.GetCode(submit.Code)
	if err != nil {
		return nil, err
	}
	submitMess := map[string]interface{}{
		"status":      submit.Result,
		"memory_cost": submit.RunningMemory,
		"time_cost":   submit.RunningTime,
		"lang":        submit.Language,
		"code":        code,
		"time":        time.Unix(submit.SubmitTime, 0).Format("2006-01-02 15:04:05"),
	}
	return submitMess, nil
}
func SubmitGetTest(SubmitId int64) (map[string]interface{}, error) {
	submit, err := models.SubmitTestGetById(SubmitId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	code, err := managers.GetCode(submit.Code)
	if err != nil {
		return nil, err
	}
	submitMess := map[string]interface{}{
		"status":      submit.Result,
		"memory_cost": submit.RunningMemory,
		"time_cost":   submit.RunningTime,
		"lang":        submit.Language,
		"code":        code,
		"output":      submit.ResultDes,
		"time":        time.Unix(submit.SubmitTime, 0).Format("2006-01-02 15:04:05"),
	}
	return submitMess, nil
}
