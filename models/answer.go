package models

import (
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
