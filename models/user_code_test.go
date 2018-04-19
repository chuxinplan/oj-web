package models

import (
	"testing"
)

func TestUserCodeCreate(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{ProblemId: 2, UserId: 4, SaveCode: "aaaaass"}
	if _, err := userCode.Create(userCode); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserCodeRemove(t *testing.T) {
	InitAllInTest()

	var userCode UserCode
	if err := userCode.Remove(3); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserCodeUpdate(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{Id: 1, ProblemId: 1, UserId: 1, SaveCode: "111111111"}
	if err := userCode.Update(userCode); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserCodeGetById(t *testing.T) {
	InitAllInTest()

	userCode := &UserCode{ProblemId: 2, UserId: 3, SaveCode: "aaaaass"}
	UserCode{}.Create(userCode)

	getUserCode, err := UserCode{}.GetById(userCode.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserCode != *userCode {
		t.Error("GetById() failed:", "%v != %v", userCode, getUserCode)
	}
}
