package models

import "time"

type Login struct {
	Id     int
	IdUser int
	Moment time.Time
}

func NewLogin(iduser int) *Login {
	return &Login{
		IdUser: iduser,
		Moment: time.Now(),
	}
}
