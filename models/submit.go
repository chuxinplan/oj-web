package models

import . "github.com/open-fightcoder/oj-web/common/store"

type Submit struct {
	Id            int64  `form:"id" json:"id"`
	ProblemId     int64  `form:"problemId" json:"problemId"`
	UserId        int64  `form:"userId" json:"userId"`
	Language      string `form:"language" json:"language"`
	SubmitTime    int64  `form:"submitTime" json:"submitTime"`
	RunningTime   int    `form:"runningTime" json:"runningTime"`
	RunningMemory int    `form:"runningMemory" json:"runningMemory"`
	Result        int    `form:"result" json:"result"`
	ResultDes     string `form:"resultDes" json:"resultDes"`
	Code          string `form:"code" json:"code"`
}

func SubmitCreate(submit *Submit) (int64, error) {
	return OrmWeb.Insert(submit)
}

func SubmitRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Submit{})
	return err
}
func SubmitUpdate(submit *Submit) error {
	_, err := OrmWeb.AllCols().ID(submit.Id).Update(submit)
	return err
}

func SubmitGetById(id int64) (*Submit, error) {
	submit := new(Submit)
	has, err := OrmWeb.Id(id).Get(submit)

	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return submit, nil
}

func SubmitGetByUserId(userId int64, currentPage int, perPage int) ([]*Submit, error) {
	submitList := make([]*Submit, 0)
	err := OrmWeb.Where("user_id=?", userId).Limit(perPage, (currentPage-1)*perPage).Find(&submitList)
	if err != nil {
		return nil, err
	}
	return submitList, nil
}

func SubmitGetByProblemId(problemId int64, currentPage int, perPage int) ([]*Submit, error) {
	submitList := make([]*Submit, 0)
	err := OrmWeb.Where("problem_id=?", problemId).Limit(perPage, (currentPage-1)*perPage).Find(&submitList)
	if err != nil {
		return nil, err
	}
	return submitList, nil
}