package managers

import (
	"github.com/open-fightcoder/oj-web/redis"
)

func RankListGet(currentPage int, perPage int) ([]string, error) {
	return redis.RankListGet(currentPage, perPage)
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
	var rankLists []map[string]interface{}
	for i, v := range userList {
		projects := make(map[string]interface{})
		projects["rank_num"] = i + 1
		projects["user_id"] = v["user_id"]
		projects["nick_name"] = 1
		projects["avator"] = "哈哈"
		projects["ac_num"] = v["ac_num"]
		rankLists = append(rankLists, projects)
	}
	return rankLists, nil
}

func GroupRankGet(currentPage int, perPage int) ([]string, error) {
	return redis.GroupRankGet(currentPage, perPage)
}
