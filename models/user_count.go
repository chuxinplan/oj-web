package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type UserCount struct {
	Id        int64
	UserId    int64
	RankNum   int
	SubmitNum int
	DateTime  int64
}

func UserCountAdd(userCount *UserCount) (int64, error) {
	return OrmWeb.Insert(userCount)
}

func UserCountRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&UserCount{})
	return err
}

func UserCountUpdate(userCount *UserCount) error {
	_, err := OrmWeb.AllCols().ID(userCount.Id).Update(userCount)
	return err
}

func UserCountGetById(id int64) (*UserCount, error) {
	userCount := new(UserCount)

	has, err := OrmWeb.Id(id).Get(userCount)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userCount, nil
}

func UserCountGetRecentMess(userId int64) ([]*UserCount, error) {
	userCountList := make([]*UserCount, 0)
	err := OrmWeb.Where("user_id=?", userId).And("date_time >= unix_timestamp(now())-2592000").Find(&userCountList)
	if err != nil {
		return nil, err
	}
	return userCountList, nil
}
