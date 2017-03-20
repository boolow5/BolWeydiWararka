package models

import (
	"time"
)

type Discussion struct {
	DiscussionId int         `json:"discussion_id" orm:"auto"`
	Guests       []*User     `json:"guests" orm:"rel(m2m)"`
	Questions    []*Question `json:"questions" orm:"rel(m2m)"`
	OpeningDate  time.Time   `json:"opening_date" orm:"auto_now_add;type(datetime)"`
	ClosingDate  time.Time   `json:"closing_date" `
}
