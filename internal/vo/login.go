package vo

import "errors"

type LoginVO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginVO) IsValid() error {

	if l.Username == "" {
		return errors.New("request username")
	}

	if l.Password == "" {
		return errors.New("request password")
	}

	return nil
}
