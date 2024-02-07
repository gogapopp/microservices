package models

import "time"

type User struct {
	Userame  string
	Email    string
	Password string
	Role     string
	CreateAt time.Time
	UpdateAt time.Time
}
