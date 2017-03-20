package models

import (
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
