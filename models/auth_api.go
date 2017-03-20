package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/boolow5/bolow/encrypt"
)

// Valid checks for empty fields in this object and returns true if all required fields are present.
func (this *User) Valid() bool {
	return (len(this.Username) > 0 && len(this.Password) > 0 && this.Profile != &Profile{})
}

// Valid checks for empty fields in this object and returns true if all required fields are present.
func (this *Profile) Valid() bool {
	return (len(this.FirstName) > 0 && len(this.LastName) > 0)
}

// SetId is used when you don't know the type of the object because you passed it as MyModel interface, then this sets the pk field of the object.
func (this *User) SetId(id int) {
	this.UserId = id
}

// SetId is used when you don't know the type of the object because you passed it as MyModel interface, then this sets the pk field of the object.
func (this *Profile) SetId(id int) {
	this.ProfileId = id
}

func (this *User) String() string {
	return this.Username
}

// Add adds new user to the database, and returns error or false if adding is failed
func (this *User) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("ValidationError: fill all the required fields")
	}

	this.CreatedAt = time.Now()
	this.UpdatedAt = time.Now()

	if saved, err := CreateUser(this); err != nil || !saved {
		return saved, err
	}

	return true, nil
}

// Update chenges the modified fields of the object and ignores the emtpy ones.
func (this *User) Update() (bool, error) {
	fmt.Println("Updating user")
	if this.UserId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &User{UserId: this.UserId}
	oldItem.Profile = this.Profile
	updated, err := UpdateItem(oldItem, this)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	fmt.Println("Updated user = ", this)
	if !updated {
		return false, nil
	}

	fmt.Println("Updated successfully")
	return true, nil
}

// Delete removes a user from the database
func (this *User) Delete() (bool, error) {
	if this.UserId < 1 {
		return false, errors.New("IntegrityError: invalid user_id")
	}

	if deleted, err := DeleteItem(this); err != nil || !deleted {
		return deleted, err
	}

	return true, nil
}

// SetPassword is the recommended way when changing password because passwords need to be hashed, they cannot be plain text.
func (this *User) SetPassword(newPassword string) error {
	var err error
	this.Password, err = encrypt.HashPassword(newPassword)
	if err != nil {
		return err
	}
	if !encrypt.IsHash(this.Password) {
		return errors.New("Failed to hash this password")
	}
	return nil
}

// Authenticate checks the username and password and returns error if the authentication fails.
func (this *User) Authenticate() (error, *User) {
	if len(this.Username) < 2 && len(this.Password) < 2 {
		return errors.New("NotEnoughDataError: Username and password are required"), this
	}

	err, user := GetUserByUsername(this.Username)
	if err != nil {
		fmt.Println(err)
		return err, this
	}
	if user.Password == "" {
		return errors.New("Invalid username or password"), this
	}

	correct, err := encrypt.AuthenticatePassword(user.Password, this.Password)

	if err != nil {
		return err, this
	}

	if !correct {
		return errors.New("Incorrect username or password"), this
	}
	user.Password = "[hidden-for-security-reasons]"
	return nil, user
}
