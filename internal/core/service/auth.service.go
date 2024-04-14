package service

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/server/auth/internal/core/models"
	"github.com/server/auth/internal/core/ports"
)

type AuthService struct {
	authDao  ports.AuthDaoPort
	loginDao ports.LoginPort
}

func (a *AuthService) Auth(mail, password string) (*map[string]interface{}, error) {
	passwordCrypto := a.md5(password)

	user, err := a.authDao.FindUserByMailAndPassword(mail, passwordCrypto)
	if err != nil {
		return nil, err
	}

	token, err := a.generateToken(strconv.Itoa(user.Id))
	if err != nil {
		return nil, err
	}

	err = a.loginDao.Save(models.NewLogin(user.Id))

	if err != nil {
		return nil, err
	}

	return &map[string]interface{}{
		"session": token,
		"user": map[string]interface{}{
			"id":   user.Id,
			"name": user.Name,
			"mail": "",
		},
	}, nil
}

func (a *AuthService) generateToken(username string) (*map[string]interface{}, error) {
	secretKey := []byte("dj25")

	expired := time.Now().Add(time.Hour * 1)

	claims := jwt.MapClaims{
		"username": username,
		"expires":  expired.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, err
	}

	return &map[string]interface{}{
		"token":   tokenString,
		"expires": expired,
	}, nil
}

func (a *AuthService) md5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func NewAuthService(authDao ports.AuthDaoPort, loginDao ports.LoginPort) *AuthService {
	return &AuthService{
		authDao:  authDao,
		loginDao: loginDao,
	}
}
