package models

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestUserCreate(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 1, UserName: strconv.FormatInt(time.Now().Unix(), 10), NickName: "11111", Avator: "11111"}
	Create(user)
	//for i := 1; i < 20; i++ {
	//user := &User{Id: 7, AccountId: 3, NickName: "fffffffffff", UserName: "aaaa"}
	//if _, err := Create(user); err != nil {
	//	t.Error("Create() failed. Error:", err)
	//}
	//}
}
func TestUserUpdate(t *testing.T) {
	InitAllInTest()

	user := &User{20, 2, "luwenjing", "哈哈哈", "女", "", "www.csdn.com",
		"www.github.com", "暂无描述", "1990-03-12", "陕西省,西安市,长安区", "大学本科", "西安邮电大学"}
	if err := Update(user); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserRemove(t *testing.T) {
	InitAllInTest()

	if err := Remove(1); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserGetById(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 1, UserName: "abcdfg", NickName: "hahaha", Description: "1111",
		Sex: "男", Birthday: "2011-10-01", DailyAddress: "西安"}
	Create(user)

	getUser, err := GetById(user.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUser != *user {
		t.Error("GetById() failed:", "%v != %v", user, getUser)
	}
}
func TestUserQueryByName(t *testing.T) {
	InitAllInTest()

	user := &User{NickName: "ssd", UserName: "rrrrrr"}
	user1 := &User{NickName: "ssd", UserName: "tttttt"}
	Create(user)
	Create(user1)

	userList, err := QueryByName("ssd")
	if err != nil {
		t.Error("QueryByName() failed:", err)
	}
	if len(userList) != 2 {
		t.Error("QueryByName() failed:", "count is wrong!")
	}
}
func TestUserGetByAccountId(t *testing.T) {
	InitAllInTest()

	user := &User{AccountId: 20}
	Create(user)

	getUser, err := GetByAccountId(1)
	if err != nil {
		t.Error("GetByAccountId() failed:", err)
	}
	fmt.Println(getUser)
	//if getAccountId != 20 {
	//	t.Error("GetByAccountId() failed:", "%v != %v", user, getUser)
	//}
}

func TestUserGetByUserName(t *testing.T) {
	InitAllInTest()

	getUser, _ := GetByUserName("luwenjing")
	fmt.Println(getUser)
}
