package models

import (
	"errors"
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

func (this *Discussion) TableName() string {
	return "discussion"
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

func (this *Discussion) Update() (bool, error) {
	fmt.Println("Updating Discussion")
	if this.DiscussionId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		fmt.Println(err_message)
		return false, errors.New(err_message)
	}

	oldItem := &Discussion{DiscussionId: this.DiscussionId}
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
