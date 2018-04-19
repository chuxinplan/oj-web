package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type Problem struct {
	Id                 int64  `form:"id" json:"id"`
	Flag               int    `form:"flag" json:"flag"`                             //1-普通题目 2-用户题目
	Status             int    `form:"status" json:"status"`                         //申请状态
	UserId             int64  `form:"userId" json:"userId"`                         //题目提供者
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
	Tag                string `form:"tag" json:"tag"`                               //题目标签
	IsSpecialJudge     bool   `form:"isSpecialJudge" json:"isSpecialJudge"`         //是否特判
	SpecialJudgeSource string `form:"specialJudgeSource" json:"specialJudgeSource"` //特判程序源代码
	Code               string `form:"code" json:"code"`                             //标准程序
}

//增加
func (this Problem) Create(problem *Problem) (int64, error) {
	_, err := OrmWeb.Insert(problem)
	if err != nil {
		return 0, err
	}
	return problem.Id, nil
}

//删除
func (this Problem) Remove(id int64) error {
	problem := new(Problem)
	_, err := OrmWeb.Id(id).Delete(problem)
	return err
}

//修改
func (this Problem) Update(problem *Problem) error {
	_, err := OrmWeb.AllCols().ID(problem.Id).Update(problem)
	return err
}

//查询
func (this Problem) GetById(id int64) (*Problem, error) {
	problem := new(Problem)
	has, err := OrmWeb.Id(id).Get(problem)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return problem, nil
}
