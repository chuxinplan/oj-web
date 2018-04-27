package managers

import (
	"github.com/open-fightcoder/oj-web/models"
	"fmt"
	"github.com/pkg/errors"
)

type Group struct {
	Id int64		//id
	Uid int64		//组长id
	Name string		//名称
	Description string		//描述
	Avator string	//头像
	Member_id []int64	//成员🆔
}



func GroupCreat(name, avator, description string, uid int64) (int64, error) {

	//判断是否存在
	group, err := models.GroupGetbyName(name)
	if err != nil {
		return 0, fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if group != nil {
		return 0, errors.New("Group is exist")
	}
	//创建
	groupinfo := &models.GroupInfo{Name:name, Uid:uid, Avator:avator, Description:description }
	createId, err := models.GroupCreate(groupinfo)
	if err != nil {
		return 0, fmt.Errorf("add group failure : %s ", err.Error())
	}

	return createId, err
}

func MemberAdd(gid, uid, owner int64) (string, error) {

	//判断组是否存在
	group, err := models.GroupGetbyId(gid)
	if err != nil {
		return "", fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if group == nil {
		return "", errors.New("Group not exist")
	}

	if group.Uid != owner {
		return "", errors.New("Permition deny")
	}

	//是否需要判断用户合法？

	//判断用户所属
	id, err:= MemberCheckByGid(uid, gid)
	if err != nil {
		return "", fmt.Errorf("get member failure : %s ", err.Error())
	}
	if id != 0 {
		return "", errors.New( "Member is already in the Group")
	}

	//添加
	groupmember := new(models.GroupMember)
	groupmember.Gid = gid
	groupmember.Uid = uid
	_, err = models.MemberAdd(groupmember)
	if err != nil {
		return "" ,fmt.Errorf("add memver failure: %s", err.Error())
	}
	return "done", nil
}


func GroupRemove(id, owner int64) error{
	//判断组是否存在
	group, err := models.GroupGetbyId(id)
	if err != nil {
		return fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if group == nil {
		return errors.New("Group not exist")
	}
	if group.Uid != owner {
		return errors.New("Permition deny")
	}

	//得到所有成员
	members, err := models.MembersQueryByGid(id)
	if err != nil {
		return fmt.Errorf("get members failure : %s ", err.Error())
	}

	//删除成员
	for _, member:= range *members {
		err = models.MemberRemove(member.Id)
		if err != nil{
			return fmt.Errorf("delete member failure : %s ", err.Error())
		}
	}

	//删除组
	err = models.GroupRemove(id)
	if err != nil {
		return fmt.Errorf("delete group failure : %s ", err.Error())
	}

	return nil
}

func MemberRemove(uid, gid, user int64) error {

	//判断组是否存在
	group, err := models.GroupGetbyId(gid)
	if err != nil {
		return fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if group == nil {
		return errors.New("Group not exist")
	}


	//判断用户是否所属组
	id, err := MemberCheckByUid(uid, gid)
	if err != nil {
		return fmt.Errorf("get member failure : %s ", err.Error())
	}
	if id == 0 {
		return errors.New("the member doesn't belong to the group")
	}

	//判断是否有权限
	if user != uid && user != group.Uid {
		return errors.New("Permition deny")
	}


	//删除
	err = models.MemberRemove(id)
	if err != nil {
		return fmt.Errorf("delete member failure : %s ", err.Error())
	}

	return nil
}


func GetGroup(id int64) (*Group, error) {
	//判断组是否存在
	groupinfo, err := models.GroupGetbyId(id)
	if err != nil {
		return nil, fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if groupinfo == nil {
		return nil, errors.New("Group not exist")
	}


	group := &Group{-1, groupinfo.Uid,groupinfo.Name, groupinfo.Description, groupinfo.Avator, nil}

	//获得成员
	members, err := models.MembersQueryByGid(id)
	if err != nil {
		return nil, fmt.Errorf("get members failure : %s ", err.Error())
	}
	if members == nil {
		return group, nil
	}

	for _, node := range *members {
		group.Member_id = append(group.Member_id, node.Uid)
	}

	return group, err

}

//判断uid成员是否加入gid组
func MemberCheckByUid(uid, gid int64) (int64, error) {
	groupmember, err := models.MembersQueryByUid(uid)
	if err != nil {
		return 0,  fmt.Errorf("get Group failure : %s ", err.Error())
	}

	for _, member := range *groupmember {
		if member.Gid == gid {
			return member.Id, nil
		}
	}

	return 0,errors.New("no Group")

}

//判断gid组内有没有uid成员
func MemberCheckByGid(uid, gid int64) (int64, error) {
	groupmember, err := models.MembersQueryByGid(gid)
	if err != nil {
		return 0,  fmt.Errorf("get Group failure : %s ", err.Error())
	}

	for _, member := range *groupmember {
		if member.Uid == uid {
			return member.Id, nil
		}
	}

	return 0,errors.New("no Group")

}

func GroupUpdate(id, owner int64, name, avator, description string) (string, error) {

	//判断组是否存在
	group, err := models.GroupGetbyId(id)
	if err != nil {
		return "", fmt.Errorf("get Group failure : %s ", err.Error())
	}
	if group == nil {
		return "", errors.New("Group not exist")
	}
	if group.Uid != owner {
		return "", errors.New("Permition deny")
	}

	if name != "" {
		group.Name = name
	}
	if avator != "" {
		group.Avator = avator
	}
	if description != "" {
		group.Description = description
	}

	err = models.GroupUpdate(group)
	if err != nil {
		return "", fmt.Errorf("get Group failure : %s ", err.Error())
	}

	return "done", err
}