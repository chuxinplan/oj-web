package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

//成员属性
type TeamMember struct {
	Id int64 	//id
	Uid int64	//用户id
	Gid int64	//组id
}


func MemberAdd(groupmember *TeamMember) (int64, error) {
	return OrmWeb.Insert(groupmember)
}

func MemberRemove(id int64)error {
	_, err := OrmWeb.Id(id).Delete(&TeamMember{})
	return err
}

func MemberGetById(id int64) (*TeamMember, error) {
	member := new(TeamMember)

	has, err := OrmWeb.Id(id).Get(member)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return member, nil

}

func MembersQueryByUid(uid int64) (*[]TeamMember, error) {
	var groupmember []TeamMember

	//不知道可不可以返回一个序列
	err := OrmWeb.Table("team_member").Select("*").
		Where("uid = ?", uid).
		Find(&groupmember)
	if err != nil {
		return nil, err
	}

	return &groupmember, err
}



func MembersQueryByGid(gid int64)(*[]TeamMember, error) {

	var groupmember []TeamMember

	//不知道可不可以返回一个序列
	err := OrmWeb.Table("team_member").Select("*").
		Where("gid = ?", gid).
		Find(&groupmember)

	if err != nil {
		return nil, err
	}

	return &groupmember, err
}

