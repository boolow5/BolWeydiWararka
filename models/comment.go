package models

import (
	"errors"
	"fmt"
	"time"
)

type Comment struct {
	CommentId int       `json:"discussion_id" orm:"auto"`
	Text      string    `json:"text" orm:"size(500)"`
	Author    *User     `json:"author" orm:"rel(fk);on_delete(cascade)"`
	Question  *Question `json:"question" orm:"rel(fk);null;on_delete(cascade)"`
	Answer    *Answer   `json:"answer" orm:"rel(fk);null;on_delete(cascade)"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Comment) Valid() bool {
	return (this.Author.UserId != 0 && this.Question.QuestionId != 0) || (this.Author.UserId != 0 && this.Answer.AnswerId != 0)
}

func (this *Comment) SetId(id int) {
	this.CommentId = id
}

func (this *Comment) String() string {
	return this.Text
}

// CRUD operations

// Add creates new comment
func (this *Comment) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("Incomplete fields")
	}
	saved, err := SaveItem(this)
	if err != nil {
		return false, err
	}
	if !saved {
		return false, errors.New("Failed to save this comment")
	}
	return true, nil
}

func (this *Comment) Update() (bool, error) {
	fmt.Println("Updating Comment")
	if this.CommentId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &Comment{CommentId: this.CommentId}
	updated, err := UpdateItem(oldItem, this)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if !updated {
		return false, nil
	}

	fmt.Println("Updated successfully")
	return true, nil
}
