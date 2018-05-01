package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
)

//成员属性
type TeamMember struct {
	Id  int64 `xorm:"pk autoincr comment('ID') BIGINT(20)"`
	Uid int64 `xorm:"not null comment('用户ID') BIGINT(20)"`
	Gid int64 `xorm:"not null comment('团队ID') BIGINT(20)"`
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
	has, err := OrmWeb.Where("uid=?", uid).Get(groupmember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return &groupmember, err
}



func MembersQueryByGid(gid int64)(*[]TeamMember, error) {

	var groupmember []TeamMember

	//不知道可不可以返回一个序列
	has, err := OrmWeb.Where("gid=?", gid).Get(&groupmember)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}

	return &groupmember, err
}

