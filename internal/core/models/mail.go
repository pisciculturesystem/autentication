package models

import "time"

type Mail struct {
	Id        int
	Mail      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (m *Mail) IsValid() error {
	return nil
}

func NewMail(mail string) *Mail {
	return &Mail{
		Mail:      mail,
		CreatedAt: time.Now(),
	}
}
