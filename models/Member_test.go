package models

import (
	"testing"
)

//func TestMemberAdd(t *testing.T) {
//	InitAllInTest()
//
//	member := &TeamMember{Uid:1, Gid:1}
//	if _, err := MemberAdd(member); err != nil {
//		t.Error("add member failed", err)
//	}
//}

func TestMemberGetById(t *testing.T) {
	InitAllInTest()

	if _, err := MemberGetById(23); err != nil {
		t.Error("get member by id failed", err)
	}
}

func TestMembersQueryByGid(t *testing.T) {
	InitAllInTest()

	_, err := MembersQueryByGid(11)
	if  err != nil {
		t.Error("query member by gid failed", err)
	}
}

func TestMembersQueryByUid(t *testing.T) {
	InitAllInTest()

	if _, err := MembersQueryByUid(23); err != nil {
		t.Error("query member by Uid failed", err)
	}
}

func TestMemberRemove(t *testing.T) {
	InitAllInTest()

	if err := MemberRemove(1); err != nil {
		t.Error("remove member woring")
	}
}

func TestMemberGetByGidUid(t *testing.T) {
	InitAllInTest()

	if _, err := MemberGetByGidUid(1, 10); err != nil {
		t.Error("remove group 1 failed", err)
	}
}