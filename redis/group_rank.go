package redis

import (
	"strconv"

	"github.com/go-redis/redis"
	. "github.com/open-fightcoder/oj-web/common/store"
)

func GroupRankAdd(groupId int64) error {
	res := RedisClient.ZAdd("group_rank", redis.Z{0, groupId})
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func GroupRankUpdate(increment int, groupId int64) error {
	res := RedisClient.ZIncrBy("group_rank", float64(increment), strconv.FormatInt(groupId, 10))
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}

func GroupRankGet(currentPage int, perPage int) ([]string, error) {
	res := RedisClient.ZRange("group_rank", int64((currentPage-1)*perPage), int64(currentPage*perPage-1))
	if res.Err() != nil {
		return nil, res.Err()
	}
	return res.Val(), nil
}
