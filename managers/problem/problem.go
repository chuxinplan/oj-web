package problem

import (
	"encoding/json"
	"math/rand"
	"time"

	"strings"

	"strconv"

	"github.com/open-fightcoder/oj-web/common/g"
	"github.com/open-fightcoder/oj-web/models"
	"github.com/open-fightcoder/oj-web/redis"
	"github.com/pkg/errors"
)

type ProblemCount struct {
	AcNum    int64 `json:"ac_num"`
	TotalNum int64 `json:"total_num"`
}

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
	problemMess := make([]map[string]interface{}, 0)
	for _, v := range problemList {
		jsonStr, err := redis.ProblemCountGet(v.Id)
		if err != nil || jsonStr == "" {
			return nil, errors.New("获取失败")
		}
		var problemCount ProblemCount
		err = json.Unmarshal([]byte(jsonStr), &problemCount)
		if err != nil {
			return nil, errors.New("获取失败")
		}
		projects := make(map[string]interface{})
		projects["id"] = v.Id
		projects["title"] = v.Title
		projects["difficulty"] = v.Difficulty
		projects["status"] = 0
		if problemCount.TotalNum == 0 {
			projects["ac_rate"] = 0.00
		} else {
			projects["ac_rate"] = strconv.FormatFloat(float64(problemCount.AcNum*100)/float64(problemCount.TotalNum), 'f', 2, 64)
		}
		problemMess = append(problemMess, projects)
	}
	result := map[string]interface{}{
		"list":         problemMess,
		"current_page": currentPage,
		"total":        count,
	}
	return result, nil
}

func ProblemGet(id int64) (map[string]interface{}, error) {
	problem, err := models.ProblemGetById(id)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	userMess, err := models.GetById(problem.UserId)
	if err != nil {
		return nil, errors.New("获取题目失败")
	}
	problemMess := map[string]interface{}{
		"id":             problem.Id,
		"user_avator":    userMess.Avator,
		"user_name":      userMess.UserName,
		"ac_rate":        0,
		"time_limit":     problem.TimeLimit,
		"memory_limit":   problem.MemoryLimit,
		"title":          problem.Title,
		"description":    problem.Description,
		"input_des":      problem.InputDes,
		"output_des":     problem.OutputDes,
		"input_case":     problem.InputCase,
		"output_case":    problem.OutputCase,
		"hint":           problem.Hint,
		"language_limit": getLimitLanguage(problem.LanguageLimit),
	}
	jsonStr, err := redis.ProblemCountGet(problem.Id)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	if jsonStr == "" {
		return problemMess, nil
	}
	var problemCount ProblemCount
	err = json.Unmarshal([]byte(jsonStr), &problemCount)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	if problemCount.TotalNum == 0 {
		problemMess["ac_rate"] = 0.00
	} else {
		problemMess["ac_rate"] = strconv.FormatFloat(float64(problemCount.AcNum*100)/float64(problemCount.TotalNum), 'f', 2, 64)
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

func getLimitLanguage(language string) []string {
	limitList := g.Conf().Common.LanguageLimit
	strs := strings.Split(language, ",")
	retList := make([]string, 0)
	for _, str := range strs {
		for _, limit := range limitList {
			if str == limit {
				retList = append(retList, str)
				break
			}
		}
	}
	return retList
}
