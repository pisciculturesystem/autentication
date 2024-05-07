package vo

import "errors"

type KeycloakValidVO struct {
	Active bool `json:"active"`
}

type KeycloakTokenVO struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type KeycloakUserVO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Mail      string `json:"mail"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (u *KeycloakUserVO) IsValid() error {
	if u.FirstName == "" {
		return errors.New("request field 'first_name'")
	}
	if u.LastName == "" {
		return errors.New("request field 'last_name'")
	}
	if u.Username == "" {
		return errors.New("request field 'username'")
	}
	if u.Mail == "" {
		return errors.New("request field 'mail'")
	}
	if u.Password == "" {
		return errors.New("request field 'password'")
	}
	return nil
}
