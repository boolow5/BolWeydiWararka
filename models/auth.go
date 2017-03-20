package models

import (
	//"gopkg.in/mgo.v2"
	"time"
)

// authentcation models
// 1. user, 2. profile
// OVERVIEW: users sign up and get access to the reader/writer privileges.
// Admin users must be put in the database by another admin

type User struct {
	UserId    int       `json:"user_id" orm:"auto"`
	Username  string    `json:"username" orm:"unique;size(30)"`
	Password  string    `json:"password" orm:"size(100)"`
	Role      string    `json:"role" orm:"size(20)"`
	Profile   *Profile  `json:"profile" orm:"rel(one);on_delete(cascade)"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

type Profile struct {
	ProfileId       int    `json:"profile_id" orm:"auto"`
	User            *User  `json:"user" orm:"reverse(one)"`
	FirstName       string `json:"first_name" orm:"size(30)"`
	MiddleName      string `json:"middle_name" orm:"size(30)"`
	LastName        string `json:"last_name" orm:"size(30)"`
	AnswerCount     int    `json:"answer_count" `
	QuestionCount   int    `json:"question_count" `
	AnswerViewCount int    `json:"anwer_view_count" `
}
