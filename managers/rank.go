package managers

import (
	"errors"

	"github.com/open-fightcoder/oj-web/redis"
)

func RankListGet(currentPage int, perPage int) ([]string, error) {
	return redis.RankListGet(currentPage, perPage)
}

func PersonRankGet(userId int64, isWeek int) ([]map[string]interface{}, error) {
	var idList []map[string]interface{}
	var err error
	if isWeek == 1 {
		idList, err = redis.PersonWeekRankGet(userId)
	} else {
		idList, err = redis.PersonMonthRankGet(userId)
	}
	if err != nil {
		return nil, errors.New("获取失败")
	}
	var rankLists []map[string]interface{}
	for i, v := range idList {
		projects := make(map[string]interface{})
		projects["rank_num"] = i + 1
		projects["user_id"] = v
		projects["nick_name"] = 1
		projects["avator"] = "哈哈"
		projects["ac_num"] = 11
		rankLists = append(rankLists, projects)
	}
	return nil, nil
}

func GroupRankGet(currentPage int, perPage int) ([]string, error) {
	return redis.GroupRankGet(currentPage, perPage)
}
