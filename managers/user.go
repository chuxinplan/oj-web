package managers

import (
	"io"

	"encoding/binary"

	"math"

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

func GetUserMess(userName string) (map[string]interface{}, error) {
	user, err := models.GetByUserName(userName)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	if user == nil {
		return nil, errors.New("用户名不存在")
	}
	problemMess := map[string]interface{}{
		"account_id":    user.AccountId,
		"user_name":     user.UserName,
		"nick_name":     user.NickName,
		"sex":           user.Sex,
		"avator":        user.Avator,
		"blog":          user.Blog,
		"git":           user.Git,
		"description":   user.Description,
		"birthday":      user.Birthday,
		"daily_address": user.DailyAddress,
		"stat_school":   user.StatSchool,
		"school_name":   user.SchoolName,
	}
	return problemMess, nil
}

func GetUserProgress(userName string) (map[string]interface{}, error) {
	user, err := models.GetByUserName(userName)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	if user == nil {
		return nil, errors.New("用户名不存在")
	}
	acNum, _ := redis.GetAcNumByUserId(user.Id)
	problemMess := map[string]interface{}{
		"pre_num":  500,
		"ac_num":   acNum,
		"fail_num": 10,
	}
	return problemMess, nil
}

func GetUserCount(userName string) (map[string]interface{}, error) {
	user, err := models.GetByUserName(userName)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	if user == nil {
		return nil, errors.New("用户名不存在")
	}
	total, err := models.SubmitCountByConds(0, user.Id, 0, "")
	if err != nil {
		return nil, errors.New("获取失败")
	}
	resMap, err := models.SubmitCountByResult(user.Id)
	if err != nil {
		return nil, errors.New("获取失败")
	}
	problemMess := make(map[string]interface{})
	problemMess["sub_num"] = total
	problemMess["all_num"] = 500
	for _, val := range resMap {
		sum := math.Float32frombits(binary.LittleEndian.Uint32(val["count"][:]))
		rate := sum / float32(total)
		switch string(val["result"][:]) {
		case "4":
			problemMess["ac_sub_num"] = sum
			problemMess["ac_rate"] = rate
			break
		case "5":
			problemMess["wa_rate"] = rate
			break
		case "6":
			problemMess["ce_rate"] = rate
			break
		case "7":
			problemMess["te_rate"] = rate
			break
		case "8":
			problemMess["me_rate"] = rate
			break
		case "9":
			problemMess["oe_rate"] = rate
			break
		case "10":
			problemMess["re_rate"] = rate
			break
		case "11":
			problemMess["se_rate"] = rate
			break
		}
	}
	//TODO 统计该UserId下所有AC且ProblemId不重复的数量
	problemMess["ac_num"] = 11
	return problemMess, nil
}
