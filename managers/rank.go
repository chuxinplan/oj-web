package managers

import (
	"github.com/open-fightcoder/oj-web/redis"
)

func RankListGet(currentPage int, perPage int) ([]string, error) {
	return redis.RankListGet(currentPage, perPage)
}

func PersonRankGet(userId int64, isWeek int) ([]string, error) {
	if isWeek == 1 {
		return redis.PersonWeekRankGet(userId)
	} else {
		return redis.PersonMonthRankGet(userId)
	}
}

func GroupRankGet(currentPage int, perPage int) ([]string, error) {
	return redis.GroupRankGet(currentPage, perPage)
}
