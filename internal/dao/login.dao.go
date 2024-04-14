package dao

import (
	"database/sql"
	"time"

	"github.com/server/auth/internal/core/models"
)

type LoginDao struct {
	db *sql.Tx
}

func (l *LoginDao) Save(model *models.Login) error {

	query := `
	   INSERT INTO api.login(iduser, moment) VALUES ($1, $2) RETURNING id;
	`
	err := l.db.QueryRow(query, model.IdUser, time.Now()).Scan(&model.Id)

	if err != nil {
		return err
	}
	return nil
}

func NewLoginDao(db *sql.Tx) *LoginDao {
	return &LoginDao{
		db: db,
	}
}
