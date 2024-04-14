package ports

import "github.com/server/auth/internal/core/models"

type MailDaoPort interface {
	FindByMail(mail string) (*models.Mail, error)
	ExistMail(mail string) bool
}
