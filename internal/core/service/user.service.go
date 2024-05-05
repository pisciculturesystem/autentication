package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/server/auth/infra/kafka"
	"github.com/server/auth/internal/core/models"
	"github.com/server/auth/internal/core/ports"
	"github.com/server/auth/internal/vo"
)

type UserService struct {
	userDao ports.UserDaoPort
	mailDao ports.MailDaoPort
}

func (a *UserService) Create(input *vo.UserCreatedVO) (int64, error) {
	if err := input.IsValid(); err != nil {
		return 0, err
	}

	exist, err := a.userDao.ExistByRegistration(input.Registration)

	if err != nil {
		return 0, err
	}

	if exist {
		errors.New("já existe um usuário cadastrado para este cpf")
	}

	exist = a.mailDao.ExistMail(input.Mail)

	if exist {
		return 0, errors.New("já existe um usuário cadastrado com este e-mail")
	}

	user := models.NewUser(input.Name, input.Registration, input.Mail, a.md5(input.Password))

	if !user.IsValid() {
		return 0, errors.New("estrutura inválida")
	}

	id, err := a.userDao.Save(user)

	if err != nil {
		return 0, err
	}

	if id > 0 {
		mail := vo.NewSendMailVO(id, input.Name, input.Mail)
		data, _ := json.Marshal(mail)
		kafka.PostMessage("sendmail", string(data))
	}

	return id, nil
}

func (a *UserService) md5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func NewUserService(userDao ports.UserDaoPort, mailDao ports.MailDaoPort) *UserService {
	return &UserService{
		userDao: userDao,
		mailDao: mailDao,
	}
}
