package models

import (
	"fmt"
	"time"
)

type Discussion struct {
	DiscussionId int         `json:"discussion_id" orm:"auto"`
	Guests       []*User     `json:"guests" orm:"rel(m2m)"`
	Questions    []*Question `json:"questions" orm:"rel(m2m)"`
	Topics       []*Topic    `json:"topics" orm:"rel(m2m)"`
	OpeningDate  time.Time   `json:"opening_date" orm:"auto_now_add;type(datetime)"`
	ClosingDate  time.Time   `json:"closing_date" `
}

func (this *Discussion) Valid() bool {
	return len(this.Guests) != 0
}

func (this *Discussion) SetId(id int) {
	this.DiscussionId = id
}

func (this *Discussion) String() string {
	return fmt.Sprintf("discussion about %v held by %v", this.Topics, this.Guests)
}
