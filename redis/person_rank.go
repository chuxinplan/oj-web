package redis

import (
	"strconv"

	"errors"

	"github.com/go-redis/redis"
	. "github.com/open-fightcoder/oj-web/common/store"
)

func PersonWeekRankAdd(userId int64) error {
	res := RedisClient.ZAdd("person_week_rank", redis.Z{0, userId})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func PersonWeekRankUpdate(increment int, userId int64) error {
	res := RedisClient.ZIncrBy("person_week_rank", float64(increment), strconv.FormatInt(userId, 10))
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func PersonWeekRankGet(userId int64) ([]map[string]interface{}, error) {
	sizeRet := RedisClient.ZCard("person_week_rank")
	if sizeRet.Err() != nil {
		return nil, errors.New("暂无数据")
	}
	size := sizeRet.Val()
	idStr := strconv.FormatInt(userId, 10)
	isExitRet := RedisClient.ZLexCount("person_week_rank", "["+idStr, "["+idStr)

	if isExitRet.Err() != nil {
		return nil, errors.New("获取失败")
	}
	if isExitRet.Val() > 0 {
		var start int64
		var end int64
		if size <= 5 {
			start = 0
			end = 4
		} else {
			res := RedisClient.ZRank("person_week_rank", idStr)
			if res.Err() != nil {
				return nil, errors.New("获取失败")
			}
			index := res.Val()
			if index < 2 {
				start = 0
			} else {
				start = index - 2
			}
			if index > size-3 {
				end = size - 2
			} else {
				end = size
			}
		}
		result := RedisClient.ZRange("person_week_rank", start, end)
		if result.Err() != nil {
			return nil, errors.New("获取失败")
		}
		var rankLists []map[string]interface{}
		for _, v := range result.Val() {
			projects := make(map[string]interface{})
			scoreRes := RedisClient.ZScore("person_week_rank", v)
			projects["user_id"] = v
			projects["ac_num"] = scoreRes
			rankLists = append(rankLists, projects)
		}
		return rankLists, nil
	} else {
		return nil, errors.New("尚未提交，暂无排名")
	}
}

func PersonMonthRankAdd(userId int64) error {
	res := RedisClient.ZAdd("person_month_rank", redis.Z{0, userId})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func PersonMonthRankUpdate(increment int, userId int64) error {
	res := RedisClient.ZIncrBy("person_month_rank", float64(increment), strconv.FormatInt(userId, 10))
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func PersonMonthRankGet(userId int64) ([]map[string]interface{}, error) {
	res := RedisClient.ZRank("person_month_rank", strconv.FormatInt(userId, 10))
	if res.Err() != nil {
		return nil, res.Err()
	}
	index := res.Val()
	start := index - 2
	if start < 0 {
		start = 0
	}
	result := RedisClient.ZRange("person_month_rank", start, index+2)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return nil, nil
}
