package models

import (
	"errors"
	"fmt"
	"time"
)

type Answer struct {
	AnswerId   int       `json:"answer_id" orm:"auto"`
	Text       string    `json:"text" orm:"size(2000)"`
	Author     *User     `json:"author" orm:"rel(fk);on_delete(cascade)"`
	Question   *Question `json:"question" orm:"rel(fk);on_delete(cascade)"`
	ViewsCount int       `json:"views_count" `
	Likes      int       `json:"likes" `
	Dislikes   int       `json:"dislikes" `

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Answer) Valid() bool {
	return (this.Question.QuestionId != 0 && len(this.Text) > 2 && this.Author.UserId != 0)
}

func (this *Answer) SetId(id int) {
	this.AnswerId = id
}

func (this *Answer) String() string {
	return fmt.Sprintf("%s's answer to %s", this.Author, this.Question)
}

// CRUD operations

// Add creates new answer
func (this *Answer) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("Incomplete fields")
	}
	saved, err := SaveItem(this)
	if err != nil {
		return false, err
	}
	if !saved {
		return false, errors.New("Failed to save this answer")
	}
	return true, nil
}

func (this *Answer) Update() (bool, error) {
	fmt.Println("Updating answer")
	if this.AnswerId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &Answer{AnswerId: this.AnswerId}
	oldItem.Question = this.Question
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
