package models

import "time"

type Session struct {
	IdUser    int
	DateTime  time.Time
	Situation string
}

func NewSession(iduser int, situation string) *Session {
	return &Session{
		IdUser:    iduser,
		DateTime:  time.Now(),
		Situation: situation,
	}
}
