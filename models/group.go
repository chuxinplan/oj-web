package models


import (
	. "github.com/open-fightcoder/oj-web/common/store"
)


//组属性
type GroupInfo struct {
	Id int64		//id
	Uid int64		//组长id
	Name string		//名称
	Description string		//描述
	Avator string	//头像
}


func GroupCreate(groupinfo *GroupInfo)(int64, error){
	return OrmWeb.Insert(groupinfo)
}

func GroupRemove(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&GroupInfo{})
	return err
}

func GroupUpdate(groupinfo *GroupInfo) error {
	_, err := OrmWeb.AllCols().ID(groupinfo.Id).Update(groupinfo)
	return err
}

func GroupGetbyId(id int64)(*GroupInfo, error) {
	groupinfo := new(GroupInfo)

	has, err := OrmWeb.Id(id).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return groupinfo, nil
}

func GroupGetbyName(name string) (*GroupInfo, error)  {

	groupinfo := new(GroupInfo)

	has, err := OrmWeb.Where("name=?", name).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return groupinfo, nil
}
