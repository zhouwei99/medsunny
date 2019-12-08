package model

import "time"

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id       int
	Uuid     string
	Email    string
	UserId   int
	CreateAt time.Time
}

func (user *User) CreateSession() (Session, error) {
	return Session{}, nil
}
