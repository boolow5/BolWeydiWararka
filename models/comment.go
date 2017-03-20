package models

import (
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
