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

func RankListGet(currentPage int, perPage int) ([]string, error) {
	res := RedisClient.ZRange("rank_list", int64((currentPage-1)*perPage), int64(currentPage*perPage-1))
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res.Val(), nil
}
