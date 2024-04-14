package vo

import "errors"

type LoginVO struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func (l *LoginVO) IsValid() error {

	if l.Mail == "" {
		return errors.New("Request mail")
	}

	if l.Password == "" {
		return errors.New("Request password")
	}

	return nil
}
