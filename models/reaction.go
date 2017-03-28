package models

import (
	"errors"
	"fmt"
	"time"
)

type Reaction struct {
	ReactionId int       `json:"reaction_id" orm:"auto"`
	Positive   bool      `json:"positive" orm:"default(true)"`
	User       *User     `json:"user" orm:"rel(fk);on_delete(cascade)"`
	Question   *Question `json:"question" orm:"null;rel(fk);on_delete(cascade)"`
	Answer     *Answer   `json:"answer" orm:"null;rel(fk);on_delete(cascade)"`
	Comment    *Comment  `json:"comment" orm:"null;rel(fk);on_delete(cascade)"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Reaction) TableName() string {
	return "reaction"
}

func (this *Reaction) Valid() bool {
	return (this.User.UserId != 0 && this.Question.QuestionId != 0) || (this.User.UserId != 0 && this.Answer.AnswerId != 0) || (this.User.UserId != 0 && this.Comment.CommentId != 0)
}

func (this *Reaction) SetId(id int) {
	this.ReactionId = id
}

func (this *Reaction) String() string {
	reaction := "liked"
	if !this.Positive {
		reaction = "disliked"
	}
	if this.Question.QuestionId > 0 {
		return fmt.Sprintf("%s %s %s", this.User, reaction, this.Question)
	}
	if this.Answer.AnswerId > 0 {
		return fmt.Sprintf("%s %s %s", this.User, reaction, this.Answer)
	}
	if this.Comment.CommentId > 0 {
		return fmt.Sprintf("%s %s %s", this.User, reaction, this.Comment)
	}
	return fmt.Sprintf("%s %s something", this.User, reaction)
}

// CRUD operations

// Add creates new reaction
func (this *Reaction) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("Incomplete fields")
	}
	saved, err := SaveItem(this)
	if err != nil {
		return false, err
	}
	if !saved {
		return false, errors.New("Failed to save this reaction")
	}
	return true, nil
}

func (this *Reaction) Update() (bool, error) {
	fmt.Println("Updating Reaction")
	if this.ReactionId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &Reaction{ReactionId: this.ReactionId}
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
