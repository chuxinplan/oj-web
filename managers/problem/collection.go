package problem

import (
	"errors"

	"github.com/open-fightcoder/oj-web/models"
)

func CollectionSet(userId int64, problemId int64, flag string) (bool, error) {
	if flag == "set" {
		userCollection := &models.UserCollection{UserId: userId, ProblemId: problemId}
		_, err := models.UserCollectionCreate(userCollection)
		if err != nil {
			return false, errors.New("收藏失败")
		}
	}
	if flag == "cancel" {
		collection, err := models.UserCollectionGetUserCollection(userId, problemId)
		if err != nil || collection == nil {
			return false, errors.New("取消收藏失败")
		}
		err = models.UserCollectionRemove(collection.Id)
		if err != nil {
			return false, errors.New("取消收藏失败")
		}
	}
	return true, nil
}

func CollectionGet(userId int64, problemId int64) (bool, error) {
	isCollection, err := models.UserCollectionGetUserCollection(userId, problemId)
	if err != nil {
		return false, errors.New("获取收藏信息失败")
	}
	if isCollection == nil {
		return false, nil
	}
	return true, nil
}
