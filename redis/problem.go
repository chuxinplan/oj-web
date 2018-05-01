package redis

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

func ProblemNumGet() (string, error) {
	res := RedisClient.Get("problem_num")
	if res.Err() != nil {
		return "", res.Err()
	}
	return res.Val(), nil
}

func ProblemNumIncr() error {
	res := RedisClient.Incr("problem_num")
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
