package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type ProblemUser struct {
	Id                 int64  `form:"id" json:"id"`
	Status             int    `form:"status" json:"status"`                         //申请状态
	UserId             int64  `form:"user_id" json:"user_id"`                       //题目提供者
	RealId             int    `form:"real_id" json:"real_id"`                       //所在真实题库的Id
	Difficulty         string `form:"difficulty" json:"difficulty"`                 //题目难度
	CaseData           string `form:"caseData" json:"caseData"`                     //测试数据
	Title              string `form:"title" json:"title"`                           //题目标题
	Description        string `form:"description" json:"description"`               //题目描述
	InputDes           string `form:"inputDes" json:"inputDes"`                     //输入描述
	OutputDes          string `form:"outputDes" json:"outputDes"`                   //输出描述
	InputCase          string `form:"inputCase" json:"inputCase"`                   //测试输入
	OutputCase         string `form:"outputCase" json:"outputCase"`                 //测试输出
	Hint               string `form:"hint" json:"hint"`                             //题目提示(可以为对样例输入输出的解释)
	TimeLimit          int    `form:"timeLimit" json:"timeLimit"`                   //时间限制
	MemoryLimit        int    `form:"memoryLimit" json:"memoryLimit"`               //内存限制
	Tag                int64  `form:"tag" json:"tag"`                               //题目标签
	IsSpecialJudge     bool   `form:"isSpecialJudge" json:"isSpecialJudge"`         //是否特判
	SpecialJudgeSource string `form:"specialJudgeSource" json:"specialJudgeSource"` //特判程序源代码
	SpecialJudgeType   string `form:"specialJudgeType" json:"specialJudgeType"`     //特判程序源代码类型
	Code               string `form:"code" json:"code"`                             //标准程序
	LanguageLimit      string `form:"languageLimit" json:"languageLimit"`           //语言限制
	Remark             string `form:"remark" json:"remark"`                         //备注
}

func ProblemUserCreate(problemUser *ProblemUser) (int64, error) {
	return OrmWeb.Insert(problemUser)
}

func ProblemUserRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&ProblemUser{})
	return err
}

func ProblemUserUpdate(problemUser *ProblemUser) error {
	_, err := OrmWeb.AllCols().ID(problemUser.Id).Update(problemUser)
	return err
}

func ProblemUserGetById(id int64) (*ProblemUser, error) {
	problemUser := new(ProblemUser)
	has, err := OrmWeb.Id(id).Get(problemUser)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problemUser, nil
}

func ProblemUserGetByUserId(userId int64, currentPage int, perPage int) ([]*ProblemUser, error) {
	problemList := make([]*ProblemUser, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&problemList)
	if err != nil {
		return nil, err
	}
	return problemList, nil
}

func ProblemUserCountByUserId(userId int64) (int64, error) {
	problemUser := &ProblemUser{}
	count, err := OrmWeb.Where("user_id=?", userId).Count(problemUser)
	if err != nil {
		return 0, err
	}
	return count, nil
}
