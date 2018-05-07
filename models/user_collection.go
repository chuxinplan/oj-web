package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type UserCollection struct {
	Id        int64 `form:"id" json:"id"`
	ProblemId int64 `form:"problem_id" json:"problem_id"` //题目ID
	UserId    int64 `form:"user_id" json:"user_id"`       //用户ID
}

func UserCollectionCreate(userCollection *UserCollection) (int64, error) {
	return OrmWeb.Insert(userCollection)
}

func UserCollectionRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&UserCollection{})
	return err
}

func UserCollectionUpdate(userCode *UserCollection) error {
	_, err := OrmWeb.AllCols().ID(userCode.Id).Update(userCode)
	return err
}

func UserCollectionGetById(id int64) (*UserCollection, error) {
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

func UserCollectionGetByUserId(userId int64, currentPage int, perPage int) ([]*UserCollection, error) {
	userCollectionList := make([]*UserCollection, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&userCollectionList)
	if err != nil {
		return nil, err
	}
	return userCollectionList, nil
}

func UserCollectionCountByUserId(userId int64) (int64, error) {
	collection := &UserCollection{}
	count, err := OrmWeb.Where("user_id=?", userId).Count(collection)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func UserCollectionGetUserCollection(userId, problemId int64) (*UserCollection, error) {
	userCollection := new(UserCollection)
	has, err := OrmWeb.Where("user_id=?", userId).And("problem_id=?", problemId).Get(userCollection)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return userCollection, nil
}

func UserCollectionGetByProblemIds(userId int64, problemId []int64) ([]*UserCollection, error) {
	session := OrmWeb.NewSession()
	if len(problemId) != 0 {
		session.In("problem_id", problemId)
	}
	userCollectionList := make([]*UserCollection, 0)

	err := session.And("user_id=?", userId).Find(&userCollectionList)
	if err != nil {
		return nil, err
	}
	return userCollectionList, nil
}
