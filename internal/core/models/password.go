package models

import "time"

type Password struct {
	Id        int
	Password  string
	EndValid  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Password) IsActive() bool {
	return true
}

func NewPassword(password string) *Password {
	return &Password{
		Password:  password,
		CreatedAt: time.Now(),
	}
}
