package vo

import "errors"

type SendMailVO struct {
	IdUser        int64  `json:"iduser"`
	NameRecipient string `json:"name_recipient"`
	MailRecipient string `json:"mail_recipient"`
}

func (s *SendMailVO) IsValid() error {
	if s.IdUser <= 0 {
		return errors.New("Request 'iduser'")
	}
	if s.MailRecipient == "" {
		return errors.New("Request 'mail_recipient'")
	}
	if s.NameRecipient == "" {
		return errors.New("Request 'name_recipient'")
	}
	return nil
}

func NewSendMailVO(iduser int64, nameRecipient string, mailRecipient string) *SendMailVO {
	return &SendMailVO{
		IdUser:        iduser,
		NameRecipient: nameRecipient,
		MailRecipient: mailRecipient,
	}
}
