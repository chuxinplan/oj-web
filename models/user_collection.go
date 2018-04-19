package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type UserCollection struct {
	Id        int64 `form:"id" json:"id"`
	ProblemId int   `form:"problemId" json:"problemId"` //题目ID
	UserId    int64 `form:"userId" json:"userId"`       //用户ID
}

func (this UserCollection) Create(userCollection *UserCollection) (int64, error) {
	_, err := OrmWeb.Insert(userCollection)
	if err != nil {
		return 0, err
	}
	return userCollection.Id, nil
}

func (this UserCollection) Remove(id int64) error {
	userCollection := new(UserCollection)
	_, err := OrmWeb.Id(id).Delete(userCollection)
	return err
}

func (this UserCollection) Update(userCode *UserCollection) error {
	_, err := OrmWeb.AllCols().ID(userCode.Id).Update(userCode)
	return err
}

func (this UserCollection) GetById(id int64) (*UserCollection, error) {
	userCollection := new(UserCollection)
	has, err := OrmWeb.Id(id).Get(userCollection)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userCollection, nil
}
