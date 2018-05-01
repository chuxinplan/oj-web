package models


import (
	. "github.com/open-fightcoder/oj-web/common/store"
)


//组属性
type Group struct {
	Id          int64  `xorm:"pk autoincr comment('团队ID') BIGINT(20)"`
	Uid         int64  `xorm:"not null comment('组长ID') BIGINT(20)"`
	Name        string `xorm:"not null comment('团队名称') unique VARCHAR(50)"`
	Description string `xorm:"not null comment('团队描述') VARCHAR(200)"`
	Avator      string `xorm:"not null comment('团队头像') VARCHAR(50)"`
}

func GroupCreate(groupinfo *Group)(int64, error){
	return OrmWeb.Insert(groupinfo)
}

func GroupRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&Group{})
	return err
}

func GroupUpdate(groupinfo *Group) error {
	_, err := OrmWeb.AllCols().ID(groupinfo.Id).Update(groupinfo)
	return err
}

func GroupGetbyId(id int64)(*Group, error) {
	groupinfo := new(Group)

	has, err := OrmWeb.Id(id).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return groupinfo, nil
}

func GroupGetbyName(name string) (*Group, error)  {

	groupinfo := new(Group)

	has, err := OrmWeb.Where("name=?", name).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return groupinfo, nil
}
