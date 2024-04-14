package ports

import (
	"github.com/server/auth/internal/core/models"
)

type UserDaoPort interface {
	Save(model *models.User) (int64, error)
	ExistByRegistration(registration string) (bool, error)
}
