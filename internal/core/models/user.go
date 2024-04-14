package models

import (
	"time"
)

type User struct {
	Id           int
	Name         string
	IsConfirmed  bool
	Registration string
	Phone        []Phone
	Mail         []Mail
	Password     []Password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) IsValid() bool {
	if u.Name == "" {
		return false
	}
	if len(u.Mail) == 0 {
		return false
	}
	if len(u.Password) == 0 {
		return false
	}
	return true
}

func NewUser(name, reigistration, mail, password string) *User {
	mails := make([]Mail, 0)
	mails = append(mails, *NewMail(mail))

	passwords := make([]Password, 0)
	passwords = append(passwords, *NewPassword(password))

	return &User{
		Name:         name,
		Mail:         mails,
		Registration: reigistration,
		Password:     passwords,
		CreatedAt:    time.Now(),
		IsConfirmed:  false,
	}
}
