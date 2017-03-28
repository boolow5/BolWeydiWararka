package models

import (
	"errors"
	"strings"
	"time"
)

type Topic struct {
	TopicId       int       `json:"topic_id" orm:"auto"`
	Text          string    `json:"text" orm:"size(2000)"`
	UniqueUrl     string    `json:"unique_url" orm:"unique"`
	ParentTopicId int       `json:"parent_topic_id" orm:"null"`
	ViewsCount    int       `json:"views_count" `
	Followers     []*User   `json:"followers" orm:"rel(m2m)"`
	CreatedAt     time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt     time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func (this *Topic) TableName() string {
	return "topic"
}

func (this *Topic) Valid() bool {
	if len(this.Text) > 0 {
		this.UniqueUrl = strings.Replace(this.Text, " ", "-", -1)
	}
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

func (this *Topic) Update() (bool, error) {
	if this.TopicId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		return false, errors.New(err_message)
	}
	this.UpdatedAt = time.Now()
	oldItem := &Topic{TopicId: this.TopicId}
	updated, err := UpdateItem(oldItem, this)
	if err != nil {
		return false, err
	}

	if !updated {
		return false, nil
	}

	return true, nil
}

func (this *Topic) Delete() (bool, error) {
	if this.TopicId < 1 {
		err_message := "ZeroIDError: give a valid id, to update this item"
		return false, errors.New(err_message)
	}

	if deleted, err := DeleteItem(this); err != nil || !deleted {
		return deleted, err
	}

	return true, nil
}
