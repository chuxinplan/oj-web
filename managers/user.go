package managers

import (
	"io"

	"github.com/open-fightcoder/oj-web/models"
	"github.com/open-fightcoder/oj-web/redis"
	"github.com/pkg/errors"
)

func UploadImage(reader io.Reader, userId int64, picType string) error {
	path, err := SaveImg(reader, userId, picType)
	if err != nil {
		return errors.New("上传失败")
	}
	user, err := models.GetById(userId)
	if err != nil || user == nil {
		return errors.New("上传失败")
	}
	user.Avator = path
	err = models.Update(user)
	if err != nil {
		return errors.New("上传失败")
	}
	return nil
}

func GetUserProgress(userId int64) (map[string]interface{}, error) {
	acNum, err := redis.GetAcNumByUserId(userId)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	problemMess := map[string]interface{}{
		"pre_num":  500,
		"ac_num":   acNum,
		"fail_num": 10,
	}
	return problemMess, nil
}
