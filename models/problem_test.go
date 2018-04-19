package models

import (
	"testing"
)

func TestProblemCreate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Flag: 10, Title: "sadas", Description: "1111", TimeLimit: 1000, MemoryLimit: 128000}
	if _, err := problem.Create(problem); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestProblemRemove(t *testing.T) {
	InitAllInTest()

	var problem Problem
	if err := problem.Remove(2); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestProblemUpdate(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Id: 1, Title: "sadas", Description: "asdasdasd"}
	if err := problem.Update(problem); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestProblemGetById(t *testing.T) {
	InitAllInTest()

	problem := &Problem{Title: "sadas", Description: "fffff"}
	Problem{}.Create(problem)

	getProblem, err := Problem{}.GetById(problem.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getProblem != *problem {
		t.Error("GetById() failed:", "%v != %v", problem, getProblem)
	}
}
