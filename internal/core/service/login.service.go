package service

import "github.com/server/auth/internal/core/ports"

type LoginService struct {
	loginDao ports.LoginPort
}

func (l *LoginService) Create(iduser int) (int, error) {

	return 0, nil
}

func NewLoginService(loginDao ports.LoginPort) *LoginService {
	return &LoginService{
		loginDao: loginDao,
	}
}
