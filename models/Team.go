package models


import (
	. "github.com/open-fightcoder/oj-web/common/store"
)


//组属性
type Team struct {
	Id          int64  `xorm:"pk autoincr comment('团队ID') BIGINT(20)"`
	Uid         int64  `xorm:"not null comment('组长ID') BIGINT(20)"`
	Name        string `xorm:"not null comment('团队名称') unique VARCHAR(50)"`
	Description string `xorm:"not null comment('团队描述') VARCHAR(200)"`
	Avator      string `xorm:"not null comment('团队头像') VARCHAR(50)"`
}

func TeamCreate(teaminfo *Team)(int64, error){
	return OrmWeb.Insert(teaminfo)
}

func TeamRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Team{})
	return err
}

func TeamUpdate(teaminfo *Team) error {
	_, err := OrmWeb.AllCols().ID(teaminfo.Id).Update(teaminfo)
	return err
}

func TeamGetbyId(id int64)(*Team, error) {
	teaminfo := new(Team)

	has, err := OrmWeb.Id(id).Get(teaminfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return teaminfo, nil
}

func TeamGetbyName(name string) (*Team, error)  {

	teaminfo := new(Team)

	has, err := OrmWeb.Where("name=?", name).Get(teaminfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return teaminfo, nil
}
