package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type UserCode struct {
	Id        int64  `form:"id" json:"id"`
	ProblemId int    `form:"problemId" json:"problemId"` //题目ID
	UserId    int64  `form:"userId" json:"userId"`       //用户ID
	SaveCode  string `form:"saveCode" json:"saveCode"`   //保存代码
}

func (this UserCode) Create(userCode *UserCode) (int64, error) {
	_, err := OrmWeb.Insert(userCode)
	if err != nil {
		return 0, err
	}
	return userCode.Id, nil
}

func (this UserCode) Remove(id int64) error {
	userCode := new(UserCode)
	_, err := OrmWeb.Id(id).Delete(userCode)
	return err
}

func (this UserCode) Update(userCode *UserCode) error {
	_, err := OrmWeb.AllCols().ID(userCode.Id).Update(userCode)
	return err
}

func (this UserCode) GetById(id int64) (*UserCode, error) {
	userCode := new(UserCode)
	has, err := OrmWeb.Id(id).Get(userCode)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userCode, nil
}
