package models

import (
	"testing"
)

func TestMemberAdd(t *testing.T) {
	InitAllInTest()

	member := &TeamMember{Uid:1, Gid:1}
	if _, err := MemberAdd(member); err != nil {
		t.Error("add member failed", err)
	}
}

func TestMemberGetById(t *testing.T) {
	InitAllInTest()

	if _, err := MemberGetById(1); err != nil {
		t.Error("get member by id failed", err)
	}
}

func TestMembersQueryByGid(t *testing.T) {
	InitAllInTest()

	if _, err := MembersQueryByGid(1); err != nil {
		t.Error("query member by gid failed", err)
	}
}

func TestMembersQueryByUid(t *testing.T) {
	InitAllInTest()

	if _, err := MembersQueryByUid(1); err != nil {
		t.Error("query member by Uid failed", err)
	}
}

func TestMemberRemove(t *testing.T) {
	InitAllInTest()

	if err := MemberRemove(1); err != nil {
		t.Error("remove member woring")
	}
}
