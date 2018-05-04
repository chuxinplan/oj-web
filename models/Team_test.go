package models

import (
	"testing"
	"fmt"
)

func TestTeamCreate(t *testing.T) {
	InitAllInTest()

	groupinfo := &Team{Avator:"touxaing", Description:"test", Uid:23, Name:"Sequinaa"}
	ret, err := TeamCreate(groupinfo);
	if  err != nil{
		t.Error("creat Team failed", err)
	}

	fmt.Println(ret)
}

func TestTeamGetbyName(t *testing.T) {

	InitAllInTest()
	ans, err := TeamGetbyName("sequinhaha")
	if err != nil {
		t.Error("get group by name failed", err)
	}
	fmt.Println(ans)
}

func TestTeamGetbyId(t *testing.T) {
	InitAllInTest()
	ans, err := TeamGetbyId(1);
	if  err != nil {
		t.Error("get group by id failed", err)
	}
	if ans == nil {
		fmt.Println("dsdsds")
	}
	fmt.Println(ans)
}

func TestTeamUpdate(t *testing.T) {
	InitAllInTest()
	groupinfo := &Team{Id:1, Avator:"", Description:"haha", Uid:002, Name:"SequinYF"}

	if err := TeamUpdate(groupinfo); err != nil {
		t.Error("update group failed", err)
	}

}

func TestTeamGetbyName2(t *testing.T) {
	InitAllInTest()
	ans, err := TeamGetbyName("SequinYF")
	if err != nil {
		t.Error("get group by name failed", err)
	}

	fmt.Println(ans)
}

func TestTeamRemove(t *testing.T) {
	InitAllInTest()
	if err := TeamRemove(1); err != nil {
		t.Error("remove group 1 failed", err)
	}
}




