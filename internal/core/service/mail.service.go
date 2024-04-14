package service

import (
	"github.com/server/auth/internal/core/ports"
)

type MailService struct {
	mailDao ports.MailDaoPort
}

func (m *MailService) IsAvailable(mail string) bool {
	return !m.mailDao.ExistMail(mail)
}

func NewMailService(mailDao ports.MailDaoPort) *MailService {
	return &MailService{
		mailDao: mailDao,
	}
}
