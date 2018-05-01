package models

import (
	"fmt"
	"testing"
	"time"
)

func TestUserCountAdd(t *testing.T) {
	InitAllInTest()

	userCount := &UserCount{7, 1, 1, 1, time.Now().Unix()}
	if _, err := UserCountAdd(userCount); err != nil {
		t.Error("Add() failed.Error:", err)
	}
}
func TestUserCountUpdate(t *testing.T) {
	InitAllInTest()

	userCount := &UserCount{1, 2, 2, 2, time.Now().Unix()}
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

	userCount := &UserCount{5, 2, 4, 1, time.Now().Unix()}
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
