package models

import "time"

type Phone struct {
	Id        int
	DDD       string
	Phone     string
	CreatedAt time.Time
	UodateAt  time.Time
}

func (p *Phone) NewPhone(ddd, phone string) *Phone {
	return &Phone{
		DDD:       ddd,
		Phone:     phone,
		CreatedAt: time.Now(),
	}
}
