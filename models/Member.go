package models

import (
	. "github.com/open-fightcoder/oj-web/common/store"
	"fmt"
)

//成员属性
type TeamMember struct {
	Id int64 	//id
	Uid int64	//用户id
	Gid int64	//组id
	Stat int
	// 1:Waiting for the group leader processing
	// 2：Wait for the user to deal with
	// 3：refused
	// 4：succeed
}



func MemberInsert(groupmember *TeamMember) (int64, error) {
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

func MemberGetByGidUid(gid, uid int64) (*TeamMember, error) {

	fmt.Println(gid, uid)
	ans := new(TeamMember)
	has, err := OrmWeb.Table("team_member").Select("*").Where("uid = ?", uid).And("gid = ?", gid).Get(ans)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return ans, nil
}

func MemberUpdate(member *TeamMember) error {
	_, err := OrmWeb.AllCols().ID(member.Id).Update(member)
	return err
}

func MembersQueryByUid(uid int64) (*[]TeamMember, error) {
	var groupmember []TeamMember

	//不知道可不可以返回一个序列
	err := OrmWeb.Table("team_member").Select("*").
		Where("uid = ?", uid).Where("stat=?", 3).
		Find(&groupmember)
	if err != nil {
		return nil, err
	}

	return &groupmember, err
}



func MembersQueryByGid(gid int64)(*[]TeamMember, error) {

	var groupmember []TeamMember

	err := OrmWeb.Table("team_member").Select("*").
		Where("gid = ?", gid).Where("stat=?", 3).
		Find(&groupmember)

	if err != nil {
		return nil, err
	}

	return &groupmember, err
}

func TeamsIDQueryByUid(id int64) (*[]TeamMember, error)  {

	var groupmember []TeamMember

	err := OrmWeb.Table("team_member").Select("*").
		Where("uid = ?", id).Where("stat = ?", 3).
		Find(&groupmember)

	if err != nil {
		return nil, err
	}

	return &groupmember, err

}
