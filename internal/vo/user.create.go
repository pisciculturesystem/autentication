package vo

import "errors"

type UserCreatedVO struct {
	Name         string `json:"name"`
	Registration string `json:"registration"`
	Mail         string `json:"mail"`
	Password     string `json:"password"`
	Phone        string `json:"phone"`
}

func (u *UserCreatedVO) IsValid() error {
	if u.Name == "" {
		return errors.New("Request field 'Name'")
	}
	if u.Registration == "" {
		return errors.New("Request field 'Registration'")
	}
	if u.Mail == "" {
		return errors.New("Request field 'Mail'")
	}
	if u.Password == "" {
		return errors.New("Request field 'Password'")
	}
	return nil
}
