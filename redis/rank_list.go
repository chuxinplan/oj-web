package redis

import (
	"strconv"

	"github.com/go-redis/redis"
	. "github.com/open-fightcoder/oj-web/common/store"
)

func RankListAdd(userId int64) error {
	res := RedisClient.ZAdd("rank_list", redis.Z{0, userId})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func RankListUpdate(increment int, userId int64) error {
	res := RedisClient.ZIncrBy("rank_list", float64(increment), strconv.FormatInt(userId, 10))
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func RankListGet(currentPage int, perPage int) ([]map[string]interface{}, error) {
	res := RedisClient.ZRevRange("rank_list", int64((currentPage-1)*perPage), int64(currentPage*perPage-1))
	if res.Err() != nil {
		return nil, res.Err()
	}
	var rankLists []map[string]interface{}
	for _, v := range res.Val() {
		projects := make(map[string]interface{})
		scoreRes := RedisClient.ZScore("rank_list", v)
		rankId := RedisClient.ZRevRank("rank_list", v)
		projects["rank_num"] = rankId.Val() + 1
		projects["user_id"] = v
		projects["ac_num"] = scoreRes.Val()
		rankLists = append(rankLists, projects)
	}
	return rankLists, nil
}
