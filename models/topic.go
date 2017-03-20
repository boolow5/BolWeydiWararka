package models

import (
	"errors"
	"time"
)

type Topic struct {
	TopicId       int       `json:"answer_id" orm:"auto"`
	Text          string    `json:"text" orm:"size(2000)"`
	ParentTopicId int       `json:"parent_topic_id" orm:"null"`
	ViewsCount    int       `json:"views_count" `
	Followers     []*User   `json:"followers" orm:"rel(m2m)"`
	CreatedAt     time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Topic) Valid() bool {
	return len(this.Text) > 1
}

func (this *Topic) SetId(id int) {
	this.TopicId = id
}

func (this *Topic) String() string {
	return this.Text
}

// CRUD operations

// Add creates new topic
func (this *Topic) Add() (bool, error) {
	if !this.Valid() {
		return false, errors.New("Incomplete fields")
	}
	saved, err := SaveItem(this)
	if err != nil {
		return false, err
	}
	if !saved {
		return false, errors.New("Failed to save this topic")
	}
	return true, nil
}
