package models

import (
	"testing"
)

func TestUserCollectionCreate(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{ProblemId: 3, UserId: 4}
	if _, err := userCollection.Create(userCollection); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserCollectionRemove(t *testing.T) {
	InitAllInTest()

	var userCollection UserCollection
	if err := userCollection.Remove(3); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserCollectionUpdate(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{Id: 1, ProblemId: 1, UserId: 1}
	if err := userCollection.Update(userCollection); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserCollectionGetById(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{ProblemId: 2, UserId: 3}
	UserCollection{}.Create(userCollection)

	getUserCollection, err := UserCollection{}.GetById(userCollection.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserCollection != *userCollection {
		t.Error("GetById() failed:", "%v != %v", userCollection, getUserCollection)
	}
}
