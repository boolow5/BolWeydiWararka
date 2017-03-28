package models

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Question struct {
	QuestionId  int         `json:"question_id" orm:"auto"`
	Text        string      `json:"text" orm:"size(200)"`
	UniqueUrl   string      `json:"unique_url" orm:"unique"`
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

func (this *Question) TableName() string {
	return "question"
}

func (this *Question) Valid() bool {
	if len(this.Text) > 0 {
		this.UniqueUrl = strings.Replace(this.Text, " ", "-", -1)
	}
	return (len(this.Text) > 2 /*&& this.Author.UserId != 0*/)
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

func (this *Question) Update() (bool, error) {
	if this.QuestionId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &Question{QuestionId: this.QuestionId}
	updated, err := UpdateItem(oldItem, this)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if !updated {
		return false, nil
	}

	return true, nil
}

func (this *Question) Delete() (bool, error) {
	if this.QuestionId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}

	if deleted, err := DeleteItem(this); err != nil || !deleted {
		return deleted, err
	}

	return true, nil
}
