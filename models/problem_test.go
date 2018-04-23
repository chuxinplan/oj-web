package models

import (
	"fmt"
	"testing"
)

func TestProblemCreate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{UserId: 4, Tag: 48, Flag: 10, Title: "sadas", Description: "1111", TimeLimit: 1000, MemoryLimit: 128000}
	if _, err := ProblemCreate(problem); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemRemove(t *testing.T) {
	InitAllInTest()

	if err := ProblemRemove(2); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUpdate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Id: 1, Title: "sadas", Description: "asdasdasd"}
	if err := ProblemUpdate(problem); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemGetById(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Title: "sadas", Description: "fffff"}
	ProblemCreate(problem)

	getProblem, err := ProblemGetById(problem.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblem != *problem {
		t.Error("GetById() failed:", problem, "!=", getProblem)
	}
}

func TestProblemGetProblem(t *testing.T) {
	InitAllInTest()
	getProblem, _ := ProblemGetProblem([]int64{}, "", "id", "asc", 1, 10)
	for i := 0; i < len(getProblem); i++ {
		fmt.Println(*getProblem[i])
	}
}

func TestProblemGetIdsByConds(t *testing.T) {
	InitAllInTest()
	getProblem, _ := ProblemGetIdsByConds([]int64{1, 2, 3}, "")
	for i := 0; i < len(getProblem); i++ {
		fmt.Println(*getProblem[i])
	}
}

func TestProblemGetByUserId(t *testing.T) {
	InitAllInTest()
	getProblem, _ := ProblemGetByUserId(1, 1, 2)
	for i := 0; i < len(getProblem); i++ {
		fmt.Println(*getProblem[i])
	}
}

func TestProblemCountByUserId(t *testing.T) {
	InitAllInTest()
	count, _ := ProblemCountByUserId(1)
	fmt.Print(count)
}

func TestProblemCountProblem(t *testing.T) {
	InitAllInTest()
	count, _ := ProblemCountProblem([]int64{1, 2, 3}, "二,三")
	fmt.Print(count)
}
