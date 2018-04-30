package managers

import (
	"io"

	"strconv"

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
	problemMess := map[string]interface{}{
		"wa_rate":    0,
		"ce_rate":    0,
		"te_rate":    0,
		"me_rate":    0,
		"oe_rate":    0,
		"re_rate":    0,
		"se_rate":    0,
		"ac_num":     11,
		"all_num":    500,
		"sub_num":    total,
		"ac_sub_num": 0,
		"ac_rate":    0,
	}

	for _, val := range resMap {
		sum, _ := strconv.ParseFloat(string(val["count"][:]), 64)
		rate := float64(sum) / float64(total)
		rateFormat := strconv.FormatFloat(rate, 'f', 2, 64)
		switch string(val["result"][:]) {
		case "4":
			problemMess["ac_sub_num"] = sum
			problemMess["ac_rate"] = rateFormat
			break
		case "5":
			problemMess["wa_rate"] = rateFormat
			break
		case "6":
			problemMess["ce_rate"] = rateFormat
			break
		case "7":
			problemMess["te_rate"] = rateFormat
			break
		case "8":
			problemMess["me_rate"] = rateFormat
			break
		case "9":
			problemMess["oe_rate"] = rateFormat
			break
		case "10":
			problemMess["re_rate"] = rateFormat
			break
		case "11":
			problemMess["se_rate"] = rateFormat
			break
		}
	}
	//TODO 统计该UserId下所有AC且ProblemId不重复的数量
	return problemMess, nil
}
