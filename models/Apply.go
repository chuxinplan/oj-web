package models

//import (
//	. "github.com/open-fightcoder/oj-web/common/store"
//)
//
//type TeamApply struct {
//	Id   int64 `xorm:"pk autoincr comment('ID') BIGINT(20)"`
//	Uid  int64 `xorm:"not null comment('用户ID') BIGINT(20)"`
//	Gid  int64 `xorm:"not null comment('组ID') BIGINT(20)"`
//	Stat int   `xorm:"not null comment('状态') INT(11)"` //0:Waiting for the group leader processing
//	// 1：Wait for the user to deal with
//	// 2：refused
//	// 3：succeed
//}
//
//
//func ApplyUpdate(apply *TeamApply) error {
//	_, err := OrmWeb.AllCols().ID(apply.Id).Update(apply)
//	return err
//}
//
//func ApplyGet(apply *TeamApply) (*TeamApply, error) {
//
//	ans := &TeamApply{}
//	has, err := OrmWeb.Where("uid = ?", apply.Uid).Where("gid = ?", apply.Gid).Get(ans)
//	if err != nil {
//		return nil, err
//	}
//	if !has {
//		return nil, nil
//	}
//	return ans, nil
//}
//
//func ApplyInsert(apply *TeamApply) (int64, error)  {
//	return OrmWeb.Insert(apply)
//}
//
//
//
//func ApplyQueryByGid(gid int64) (*[]TeamApply, error) {
//	var applys []TeamApply
//
//	err := OrmWeb.Table("team_apply").Select("*").
//		Where("gid = ?", gid).
//		Find(&applys)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return &applys, err
//}