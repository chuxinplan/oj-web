package managers

import (
	"io"

	"github.com/open-fightcoder/oj-web/models"
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
	problemMess := map[string]interface{}{
		"totle_num": 500,
		"ac_num":    10,
		"fail_num":  10,
	}
	return problemMess, nil
}
