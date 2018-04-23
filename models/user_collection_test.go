package models

import (
	"testing"
)

func TestUserCollectionCreate(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{ProblemId: 3, UserId: 4}
	if _, err := UserCollectionCreate(userCollection); err != nil {
		t.Error("Create() failed. Error:", err)
	}
}
func TestUserCollectionRemove(t *testing.T) {
	InitAllInTest()

	if err := UserCollectionRemove(3); err != nil {
		t.Error("Remove() failed. Error:", err)
	}
}
func TestUserCollectionUpdate(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{Id: 1, ProblemId: 1, UserId: 1}
	if err := UserCollectionUpdate(userCollection); err != nil {
		t.Error("Update() failed. Error:", err)
	}
}
func TestUserCollectionGetById(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{ProblemId: 2, UserId: 3}
	UserCollectionCreate(userCollection)

	getUserCollection, err := UserCollectionGetById(userCollection.Id)
	if err != nil {
		t.Error("GetById() failed:", err.Error())
	}

	if *getUserCollection != *userCollection {
		t.Error("GetById() failed:", userCollection, "!=", getUserCollection)
	}
}

func TestUserCollectionGetUserCollection(t *testing.T) {
	InitAllInTest()

	userCollection := &UserCollection{ProblemId: 5, UserId: 6}
	UserCollectionCreate(userCollection)

	getUserCollection, err := UserCollectionGetUserCollection(userCollection.UserId, userCollection.ProblemId)
	if err != nil {
		t.Error("GetUserCollection() failed:", err.Error())
	}

	if *getUserCollection != *userCollection {
		t.Error("GetUserCollection() failed:", userCollection, "!=", getUserCollection)
	}
}
