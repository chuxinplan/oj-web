package problem

import (
	"errors"
	"strconv"
	"strings"

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

func CollectionGet(userId int64, problemId string) ([]bool, error) {
	ids := []int64{}
	if problemId != "" {
		strs := strings.Split(problemId, ",")
		for i := 0; i < len(strs); i++ {
			id, _ := strconv.ParseInt(strs[i], 10, 64)
			ids = append(ids, id)
		}
	}
	isCollection, err := models.UserCollectionGetByProblemIds(userId, ids)
	if err != nil {
		return nil, errors.New("获取收藏信息失败")
	}
	resMap := []bool{}
	for i := 0; i < len(ids); i++ {
		resMap = append(resMap, false)
	}
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(isCollection); j++ {
			if ids[i] == isCollection[j].ProblemId {
				resMap[i] = true
				break
			}
		}
	}
	return resMap, nil
}
