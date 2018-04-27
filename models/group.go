package models


import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

type Group struct {
	Info *GroupInfo
	Memeber *[]GroupMember
}

//组属性
type GroupInfo struct {
	Id int64		//id
	Name string		//名称
	Description string		//描述
	Avator string	//头像
}

//成员属性
type GroupMember struct {
	Id int64	//用户id
	Gid int64	//组id
}

func GroupCreate(groupinfo *GroupInfo)(int64, error){
	return OrmWeb.Insert(groupinfo)
}

func GroupAdd(groupmember *GroupMember) (int64, error) {
	return OrmWeb.Insert(groupmember)
}

func GroupDestroy(id int64) error {
	_, err := OrmWeb.Id(id).Delete(&GroupInfo{})
	return err
}

func GroupRemove(id int64)error {
	_, err := OrmWeb.Id(id).Delete(&GroupMember{})
	return err
}

func GroupUpdate(groupinfo *GroupInfo) error {
	_, err := OrmWeb.AllCols().ID(groupinfo.Id).Update(groupinfo)
	return err
}

func MemberUpdate(member *GroupMember) error {
	_, err:= OrmWeb.AllCols().ID(member.Id).Update(member)
	return err
}

func GroupGetbyId(id int64)(*Group, error) {
	group := new(Group)
	groupinfo := new(GroupInfo)
	groupmember := new([]GroupMember)

	has, err := OrmWeb.Id(id).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	has, err = OrmWeb.Where("gid=?", groupinfo.Id).Get(groupmember)
	if err != nil {
		return nil, err
	}

	group.Memeber = groupmember
	group.Info = groupinfo

	return group, nil
}

func GroupGetbyName(name string) (*Group, error)  {

	group := new(Group)
	groupinfo := new(GroupInfo)
	groupmember := new([]GroupMember)

	has, err := OrmWeb.Where("name=?", name).Get(groupinfo)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	has, err = OrmWeb.Where("gid=?", groupinfo.Id).Get(groupmember)

	if err != nil {
		return nil, err
	}

	group.Memeber = groupmember
	group.Info = groupinfo

	return group, nil
}
