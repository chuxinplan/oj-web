package models

import (
	"testing"
	"fmt"
)

func TestGroupCreate(t *testing.T) {
	InitAllInTest()

	groupinfo := &Group{Avator:"touxaing", Description:"test", Uid:001, Name:"Sequin"}
	if _, err := GroupCreate(groupinfo); err != nil{
		t.Error("creat Group failed", err)
	}
}

func TestGroupGetbyName(t *testing.T) {

	InitAllInTest()
	ans, err := GroupGetbyName("Sequin")
	if err != nil {
		t.Error("get group by name failed", err)
	}
	fmt.Println(ans)
}

func TestGroupGetbyId(t *testing.T) {
	InitAllInTest()
	if _, err := GroupGetbyId(1); err != nil {
		t.Error("get group by id failed", err)
	}
}

func TestGroupUpdate(t *testing.T) {
	InitAllInTest()
	groupinfo := &Group{Avator:"", Description:"haha", Uid:002, Name:"SequinYF"}

	if err := GroupUpdate(groupinfo); err != nil {
		t.Error("update group failed", err)
	}
}

func TestGroupGetbyName2(t *testing.T) {
	InitAllInTest()
	ans, err := GroupGetbyName("SequinYF")
	if err != nil {
		t.Error("get group by name failed", err)
	}

	fmt.Println(ans)
}

func TestGroupRemove(t *testing.T) {
	InitAllInTest()
	if err := GroupRemove(1); err != nil {
		t.Error("remove group 1 failed", err)
	}
}


