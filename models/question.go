package models

import (
	"errors"
	"time"
)

type Question struct {
	QuestionId  int         `json:"question_id" orm:"auto"`
	Text        string      `json:"text" orm:"size(200)"`
	Description string      `json:"description" orm:"size(500)"`
	Author      *User       `json:"author" orm:"rel(fk);on_delete(cascade)"`
	Discussion  *Discussion `json:"discussion" orm:"rel(fk);null;on_delete(set_null)"`
	ViewsCount  int         `json:"views_count" `
	Likes       int         `json:"likes" `
	Dislikes    int         `json:"dislikes" `
	Answers     []*Answer   `json:"answers" orm:"reverse(many)"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Question) Valid() bool {
	return (len(this.Text) > 2 && this.Author.UserId != 0)
}

func (this *Question) SetId(id int) {
	this.QuestionId = id
}

func (this *Question) String() string {
	return this.Text
}

// CRUD operations

// Add creates new question
func (this *Question) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("Incomplete fields")
	}
	saved, err := SaveItem(this)
	if err != nil {
		return false, err
	}
	if !saved {
		return false, errors.New("Failed to save this question")
	}
	return true, nil
}
