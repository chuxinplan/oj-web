package models

import (
	"testing"
)

func TestUserCodeCreate(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{ProblemId: 2, UserId: 4, SaveCode: "aaaaass"}
	if _, err := UserCodeCreate(userCode); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserCodeRemove(t *testing.T) {
	InitAllInTest()

	if err := UserCodeRemove(3); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserCodeUpdate(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{Id: 1, ProblemId: 1, UserId: 1, SaveCode: "111111111"}
	if err := UserCodeUpdate(userCode); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserCodeGetById(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{ProblemId: 2, UserId: 3, SaveCode: "aaaaass"}
	UserCodeCreate(userCode)

	getUserCode, err := UserCodeGetById(userCode.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserCode != *userCode {
		t.Error("GetById() failed:", userCode, "!=", getUserCode)
	}
}
func TestUserCodeGetUserCode(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{ProblemId: 5, UserId: 5, SaveCode: "aaaaass"}
	UserCodeCreate(userCode)

	getUserCollection, err := UserCodeGetUserCode(userCode.UserId, userCode.ProblemId)
	if err != nil {
		t.Error("GetUserCollection() failed:", err.Error())
	}

	if *getUserCollection != *userCode {
		t.Error("GetUserCollection() failed:", userCode, "!=", getUserCollection)
	}
}
