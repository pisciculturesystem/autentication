package dao

import (
	"database/sql"
	"encoding/json"

	"github.com/server/auth/internal/core/models"
)

type UserDao struct {
	db *sql.DB
}

func (u *UserDao) Save(model *models.User) (int64, error) {

	tx, err := u.db.Begin()
	if err != nil {
		return 0, err
	}

	configurations, err := json.Marshal(models.NewConfigurationDetail())
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var idconfiguration int64

	sql := "INSERT INTO api.configuration(configuration) VALUES($1) RETURNING id"
	err = u.db.QueryRow(sql, configurations).Scan(&idconfiguration)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int64

	sql = "INSERT INTO api.user(name, registration, idconfiguration) VALUES ($1, $2, $3) RETURNING id"
	err = u.db.QueryRow(sql, model.Name, model.Registration, idconfiguration).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	sql = "INSERT INTO api.mail(iduser, mail) VALUES($1, $2)"
	_, err = u.db.Exec(sql, id, model.Mail[0].Mail)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	sql = "INSERT INTO api.password(iduser, password) VALUES($1, $2)"
	_, err = u.db.Exec(sql, id, model.Password[0].Password)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *UserDao) ExistByRegistration(registration string) (bool, error) {
	return false, nil
}

func NewUserDao(db *sql.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}
