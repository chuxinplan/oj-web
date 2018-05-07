package models

import (
	"fmt"
	"testing"
)

func TestProblemUserCreate(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{UserId: 25, Tag: 10, Title: "四", Description: "四"}
	if _, err := ProblemUserCreate(problemUser); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemUserRemove(t *testing.T) {
	InitAllInTest()

	if err := ProblemUserRemove(2); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUserUpdate(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{Id: 1, Title: "sadas", Description: "asdasdasd"}
	if err := ProblemUserUpdate(problemUser); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemUserGetById(t *testing.T) {
	InitAllInTest()

	problemUser := &ProblemUser{Title: "sadas", Description: "fffff"}
	ProblemUserCreate(problemUser)

	getProblemUser, err := ProblemUserGetById(problemUser.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblemUser != *problemUser {
		t.Error("GetById() failed:", problemUser, "!=", getProblemUser)
	}
}

func TestProblemUserGetByUserId(t *testing.T) {
	InitAllInTest()
	getProblemUser, _ := ProblemUserGetByUserId(1, 1, 2)
	for i := 0; i < len(getProblemUser); i++ {
		fmt.Println(*getProblemUser[i])
	}
}

func TestProblemUserCountByUserId(t *testing.T) {
	InitAllInTest()
	count, _ := ProblemUserCountByUserId(1)
	fmt.Print(count)
}
