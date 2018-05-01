package models

import (
	"fmt"
	"testing"
)

func TestUserCountAdd(t *testing.T) {
	InitAllInTest()

	userCount := &UserCount{4, 1, 1, 1, "2018-04-20"}
	if _, err := UserCountAdd(userCount); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}
func TestUserCountUpdate(t *testing.T) {
	InitAllInTest()

	userCount := &UserCount{1, 2, 2, 2, "2016-01-02"}
	if err := UserCountUpdate(userCount); err != nil {
		t.Error("Update() failed.Error:", err)
	}
}
func TestUserCountRemove(t *testing.T) {
	InitAllInTest()

	if err := UserCountRemove(1); err != nil {
		t.Error("Remove() failed.Error:", err)
	}
}
func TestUserCountGetById(t *testing.T) {
	InitAllInTest()

	userCount := &UserCount{5, 2, 4, 1, "2016-01-02"}
	UserCountAdd(userCount)

	getUserCount, err := UserCountGetById(userCount.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserCount != *userCount {
		t.Error("GetById() failed:", userCount, "!=", getUserCount)
	}
}

func TestUserCountGetRecentMess(t *testing.T) {
	InitAllInTest()

	userCount, _ := UserCountGetRecentMess(1)
	for _, v := range userCount {
		fmt.Println(v)
	}
}
