package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

//成员属性
type GroupMember struct {
	Id int64 	//id
	Uid int64	//用户id
	Gid int64	//组id
}


func MemberAdd(groupmember *GroupMember) (int64, error) {
	return OrmWeb.Insert(groupmember)
}

func MemberRemove(id int64)error {
	_, err := OrmWeb.Id(id).Delete(&GroupMember{})
	return err
}

func MemberGetById(id int64) (*GroupMember, error) {
	member := new(GroupMember)

	has, err := OrmWeb.Id(id).Get(member)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return member, nil

}

func MembersQueryByUid(uid int64) (*[]GroupMember, error) {
	var groupmember []GroupMember

	//不知道可不可以返回一个序列
	has, err := OrmWeb.Where("uid=?", uid).Get(groupmember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return &groupmember, err
}



func MembersQueryByGid(gid int64)(*[]GroupMember, error) {

	var groupmember []GroupMember

	//不知道可不可以返回一个序列
	has, err := OrmWeb.Where("gid=?", gid).Get(groupmember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return &groupmember, err
}

