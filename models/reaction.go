package models

import (
	"time"
)

type Reaction struct {
	Positive bool      `json:"positive" orm:"default(true)"`
	User     *User     `json:"user" orm:"rel(fk);on_delete(cascade)"`
	Question *Question `json:"question" orm:"null;rel(fk);on_delete(cascade)"`
	Answer   *Answer   `json:"answer" orm:"null;rel(fk);on_delete(cascade)"`
	Comment  *Comment  `json:"comment" orm:"null;rel(fk);on_delete(cascade)"`

	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}
