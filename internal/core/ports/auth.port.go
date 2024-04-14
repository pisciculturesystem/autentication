package ports

import "github.com/server/auth/internal/core/models"

type AuthDaoPort interface {
	FindUserByMailAndPassword(mail, password string) (*models.User, error)
}
