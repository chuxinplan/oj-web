package models

import (
	"fmt"
	"testing"
)

func TestSubmitCreate(t *testing.T) {
	InitAllInTest()

	submit := Submit{ProblemId: 3, UserId: 1, Language: "GO", SubmitTime: 123456, RunningTime: 12, RunningMemory: 12, ResultDes: "123456"}

	if _, err := SubmitCreate(&submit); err != nil {
		t.Error("create submit error")
	}

}

func TestSubmitRemove(t *testing.T) {
	InitAllInTest()

	if SubmitRemove(1) != nil {
		t.Error("submit remove error")
	}
}

func TestSubmitUpdate(t *testing.T) {
	InitAllInTest()
	submit := Submit{ProblemId: 1, UserId: 1, Language: "GO", SubmitTime: 123456, RunningTime: 12, RunningMemory: 12, ResultDes: "654321"}
	if SubmitUpdate(&submit) != nil {
		t.Error("update submit error")
	}
}

func TestSubmitGetById(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetById(1); err != nil {
		t.Error("get submit by id error")
	}
}

func TestSubmitGetByUserId(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetByUserId(1, 1, 1); err != nil {
		t.Error("get submit by user id error")
	}
}

func TestSubmitGetByProblemId(t *testing.T) {
	InitAllInTest()

	if _, err := SubmitGetByProblemId(1, 1, 1); err != nil {
		t.Error("get submit by problem id error")
	}
}

func TestSubmitGetByConds(t *testing.T) {
	InitAllInTest()

	submit, _ := SubmitGetByConds(0, 0, 0, "", 2, 2)
	for _, v := range submit {
		fmt.Println(v)
	}
}
