package dao

import (
	"database/sql"
	"errors"

	"github.com/server/auth/internal/core/models"
)

type AuthDao struct {
	db *sql.DB
}

func (a *AuthDao) FindUserByMailAndPassword(mail, password string) (*models.User, error) {

	query := `
	 SELECT u.id, 
	        u.name
	   FROM api.user u 
	   JOIN api.mail m ON m.iduser = u.id
	   JOIN api.password p ON p.iduser = u.id
	  WHERE m.mail = $1 
	    AND m.endvalid IS NULL 
	    AND p.password = $2
		AND p.endvalid IS NULL
	    AND u.confirmed = true`

	row := a.db.QueryRow(query, mail, password)

	var user models.User

	err := row.Scan(&user.Id, &user.Name) //, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, errors.New("User not found")
	}

	return &user, nil
}

func NewAuthDao(db *sql.DB) *AuthDao {
	return &AuthDao{
		db: db,
	}
}
