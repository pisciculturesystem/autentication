package dao

import (
	"database/sql"

	"github.com/server/auth/internal/core/models"
)

type MailDao struct {
	db *sql.DB
}

func (m *MailDao) ExistMail(mail string) bool {
	query := `
     SELECT m.id 
	   FROM api.mail m 
      WHERE UPPER(m.mail) = UPPER($1) 
	    AND endvalid IS NULL`

	var idmail int
	err := m.db.QueryRow(query, mail).Scan(&idmail)
	return err != sql.ErrNoRows
}

func (m *MailDao) FindByMail(mail string) (*models.Mail, error) {
	query := `
	  SELECT m.id, 
	         m.mail, 
		  	 m.createdat 
	    FROM api.mail m 
   	   WHERE UPPER(m.mail) = UPPER($1) 
	     AND endvalid IS NULL`

	row := m.db.QueryRow(query, mail)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var response models.Mail
	row.Scan(&response.Id, &response.Mail, &response.CreatedAt)

	return &response, nil
}

func NewMailDao(db *sql.DB) *MailDao {
	return &MailDao{
		db: db,
	}
}
