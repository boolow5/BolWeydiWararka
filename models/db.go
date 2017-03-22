package models

// This file is used to abstract a lot of database manipulations, to ease future migration to other databases.
// All function for adding, modifiying and removing items from the database are all in this file.
// Currently this file uses Sqlite as database nd Beego orm as a layer for accessing the database

import (
	"errors"
)

// SaveItem saves any object that implements MyModel interface.
func SaveItem(item MyModel) (bool, error) {
	o.Begin()
	rows_affected, err := o.Insert(item)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No rows inserted")
	}
	o.Commit()
	return true, nil
}

// UpdateItem modifies the non-empty fields in the newItem
func UpdateItem(oldItem, newItem MyModel) (bool, error) {
	o.Begin()
	err := o.Read(oldItem)
	if err != nil {
		return false, err
	}
	rows_affected, err := o.Update(newItem)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No rows affected")
	}
	err = o.Commit()
	if err != nil {
		return false, err
	}

	return true, nil
}

// DeleteItem removes items from the database and returns false and/or error if it fails
func DeleteItem(item MyModel) (bool, error) {
	o.Begin()
	rows_affected, err := o.Delete(item)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		return false, errors.New("0 Rows deleted")
	}
	o.Commit()
	return true, nil
}

// GetItems fetchs an array of items from the database
func GetItems(itemArray []MyModel) []MyModel {
	return itemArray
}

func GetQuestionById(item *Question) (*Question, error) {
	if item.QuestionId < 1 {
		return item, errors.New("Invalid question id")
	}
	err := o.Read(item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// GetOneItem fetchs a single item from the database
func GetOneItem(item MyModel) (MyModel, error) {
	err := o.Read(item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// GetUserById finds a user object by its primary key
func GetUserById(id int) (error, *User) {
	user := User{UserId: id}
	err := o.Read(&user)
	if err != nil {
		return err, &user
	}
	return nil, &user
}

// GetUserByUsername fetchs a user by its username field
func GetUserByUsername(username string) (error, *User) {
	user := User{Username: username}
	err := o.Read(&user, "username")
	if err != nil {
		return err, &user
	}
	return nil, &user
}

// CreateUser adds new user to the database
func CreateUser(item *User) (bool, error) {
	o.Begin()
	rows_affected, err := o.Insert(item.Profile)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No profile inserted")
	}
	rows_affected, err = o.Insert(item)
	if err != nil {
		o.Rollback()
		return false, err
	}
	if rows_affected < 1 {
		o.Rollback()
		return false, errors.New("No user inserted")
	}
	o.Commit()
	return true, nil
}
