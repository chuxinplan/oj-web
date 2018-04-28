package problem

import (
	"math/rand"
	"time"

	"github.com/open-fightcoder/oj-web/models"
	"github.com/pkg/errors"
)

func ProblemList(origin string, tag string, sort int, isAsc int, currentPage int, perPage int) (map[string]interface{}, error) {
	//TODO 排序条件 1-编号 2-难度 3-通过率
	sortKey := "id"
	isAscKey := "asc"
	if isAsc == 2 {
		isAscKey = "desc"
	}
	problemList, err := models.ProblemGetProblem(origin, tag, sortKey, isAscKey, currentPage, perPage)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	count, err := models.ProblemCountProblem(origin, tag)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	problemMess := map[string]interface{}{
		"list":         problemList,
		"current_page": currentPage,
		"total":        count,
	}
	return problemMess, nil
}

func ProblemGet(id int64) (map[string]interface{}, error) {
	problem, err := models.ProblemGetById(id)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	//TODO 从Redis中去获取ac_rate
	userMess, err := models.GetById(problem.UserId)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	problemMess := map[string]interface{}{
		"id":           problem.Id,
		"user_id":      problem.UserId,
		"nick_name":    userMess.NickName,
		"ac_rate":      11,
		"time_limit":   problem.TimeLimit,
		"memory_limit": problem.MemoryLimit,
		"title":        problem.Title,
		"description":  problem.Description,
		"input_des":    problem.InputDes,
		"output_des":   problem.OutputDes,
		"input_case":   problem.InputCase,
		"output_case":  problem.OutputCase,
		"hint":         problem.Hint,
	}
	return problemMess, nil
}

func ProblemRandom(origin string, tag string) (map[string]interface{}, error) {
	problemList, err := models.ProblemGetIdsByConds(origin, tag)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	size := len(problemList)
	ids := []int64{}
	for i := 0; i < size; i++ {
		ids = append(ids, problemList[i].Id)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return ProblemGet(ids[r.Intn(size)])
}
