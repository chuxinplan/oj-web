package managers

import (
	"encoding/json"
	"strconv"

	"github.com/open-fightcoder/oj-web/models"
	"github.com/open-fightcoder/oj-web/redis"
	"github.com/pkg/errors"
)

func RankListGet(currentPage int, perPage int) (map[string]interface{}, error) {
	rankList, err := redis.RankListGet(currentPage, perPage)
	if err != nil {
		return nil, err
	}
	rankLists := make([]map[string]interface{}, 0)
	for _, v := range rankList {
		userId, _ := strconv.ParseInt(v["user_id"].(string), 10, 64)
		user, err := models.GetById(userId)
		if err != nil {
			return nil, errors.New("获取失败")
		}
		if user == nil {
			return nil, errors.New("用户不存在!")
		}
		jsonStr, err := redis.SubmitCountGet(user.Id)
		if err != nil {
			return nil, errors.New("获取失败")
		}
		var submitCount SubmitCount
		err = json.Unmarshal([]byte(jsonStr), &submitCount)
		if err != nil {
			return nil, errors.New("获取失败")
		}
		projects := make(map[string]interface{})
		projects["rank_num"] = v["rank_num"]
		projects["user_name"] = user.UserName
		projects["nick_name"] = user.NickName
		projects["avator"] = user.Avator
		projects["ac_num"] = v["ac_num"]
		projects["total_num"] = submitCount.WrongAnswer + submitCount.CompilationError + submitCount.TimeLimitExceeded +
			submitCount.MemoryLimitExceeded + submitCount.OutputLimitExceeded + submitCount.RuntimeError + submitCount.SystemError + submitCount.Accepted
		rankLists = append(rankLists, projects)
	}
	rankMess := map[string]interface{}{
		"list":         rankLists,
		"current_page": currentPage,
		"total":        redis.RankListCount(),
	}
	return rankMess, nil
}

func PersonRankGet(userId int64, isWeek int) ([]map[string]interface{}, error) {
	var userList []map[string]interface{}
	var err error
	if isWeek == 1 {
		userList, err = redis.PersonWeekRankGet(userId)
	} else {
		userList, err = redis.PersonMonthRankGet(userId)
	}
	if err != nil {
		return nil, err
	}
	rankLists := make([]map[string]interface{}, 0)
	for _, v := range userList {
		userId, _ := strconv.ParseInt(v["user_id"].(string), 10, 64)
		user, err := models.GetById(userId)
		if err != nil {
			return nil, errors.New("获取失败")
		}
		projects := make(map[string]interface{})
		projects["rank_num"] = v["rank_num"]
		projects["user_id"] = v["user_id"]
		projects["nick_name"] = user.NickName
		projects["user_name"] = user.UserName
		projects["avator"] = user.Avator
		projects["ac_num"] = v["ac_num"]
		rankLists = append(rankLists, projects)
	}
	return rankLists, nil
}

func GroupRankGet(currentPage int, perPage int) ([]map[string]interface{}, error) {
	return redis.GroupRankGet(currentPage, perPage)
}
