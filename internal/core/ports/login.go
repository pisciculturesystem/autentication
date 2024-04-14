package ports

import "github.com/server/auth/internal/core/models"

type LoginPort interface {
	Save(model *models.Login) error
}
