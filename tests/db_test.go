package tests

import (
	"testing"

	"github.com/boolow5/BolWeydi/models"
)

// This file is used to abstract a lot of database manipulations, to ease future migration to other databases.
// All function for adding, modifiying and removing items from the database are all in this file.
// Currently this file uses Sqlite as database nd Beego orm as a layer for accessing the database

func TestSaveItem(t *testing.T) {
	// 1. check success
	// create item
	q := models.Question{Text: "Is this working?", Author: &models.User{UserId: 1}}
	// save the item
	saved, err := models.SaveItem(&q)
	// check if the saving succeeded
	if err != nil {
		t.Fatalf("Saving Failed: %v", err)
	}
	if !saved {
		t.Fatalf("Saving Failed, without errors")
	}
	// 2. check failure
	// create item
	a := models.Answer{}
	// save the item
	saved, err = models.SaveItem(&a)
	// check if the saving fails
	if err == nil {
		t.Fatalf("Saving should fail because item is null")
	}
	if saved {
		t.Fatalf("Saving should not succeed this time")
	}
}
