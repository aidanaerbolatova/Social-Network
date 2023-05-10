package models

import (
	"html/template"
	"time"
)

type User struct {
	Id       int
	Email    string
	Username string
	Password string
	Method   string
}

type Token struct {
	Id        int
	UserId    int
	AuthToken string
	ExpiresAT time.Time
}

type Post struct {
	Id         int
	UserId     int
	Title      string
	Text       string
	Categories string
	CreatedAt  string
	Author     string
	Like       string
	Dislike    string
	Image      template.URL
}

type MyPost struct {
	User         string
	Posts        []Post
	Notification []Notifications
}

type PostID struct {
	Username string
	Post     Post
}

type CommentID struct {
	Username string
	Comment  Comments
}

type UserCategory struct {
	Username string
	Category []Post
}

type Comments struct {
	Id        int
	UserId    int
	PostId    int
	Comment   string
	CreatedAt string
	Author    string
	Like      string
	Dislike   string
}

type GetComments struct {
	User     string
	Post     Post
	Comments []Comments
	Images   []template.URL
}

type Evaluate struct {
	Id     int
	PostId int
	UserId int
	Vote   int
}
type Vote struct {
	Like    string
	Dislike string
}

type EvaluateComment struct {
	Id        int
	UserId    int
	CommentId int
	Vote      int
}

type Notifications struct {
	From   string
	To     string
	Action string
}

type ErrorHTTP struct {
	Status  int
	Message error
}

type Follow struct {
	UserId   int
	AuthorId int
}

type AllFollow struct {
	User      User
	Followers []User
	Following []User
}
