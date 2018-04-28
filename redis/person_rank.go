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
	//判断zset中是否存在该userId
	idStr := strconv.FormatInt(userId, 10)
	isExitRet := RedisClient.ZLexCount("person_week_rank", "["+idStr, "["+idStr)
	if isExitRet.Err() != nil {
		return nil, errors.New("获取失败")
	}
	if isExitRet.Val() > 0 {
		//存在，获取zset下标
		//取5条记录，包括userId和AC数
	} else {
		return nil, errors.New("尚未提交，暂无排名")
	}

	res := RedisClient.ZRank("person_week_rank", strconv.FormatInt(userId, 10))
	//错误处理:在Redis ZSet中不存在UserId的情况
	if res.Err() != nil {
		return nil, res.Err()
	}
	index := res.Val()
	start := index - 2
	if start < 0 {
		start = 0
	}
	result := RedisClient.ZRange("person_week_rank", start, index+2)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result.Val(), nil
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

func PersonMonthRankGet(userId int64) ([]string, error) {
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
	return result.Val(), nil
}
