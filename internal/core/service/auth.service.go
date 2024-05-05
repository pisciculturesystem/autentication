package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/server/auth/internal/vo"
)

type AuthService struct {
}

func (a *AuthService) Auth(mail, password string) (*map[string]interface{}, error) {

	var form = url.Values{
		"client_id":     {"pisciculsoft"},
		"grant_type":    {"password"},
		"username":      {mail},
		"password":      {password},
		"client_secret": {"YJKvL0QvS7NL7h63Cjf5T9Y4v2tcdw39"},
	}
	var url = "http://keycloak:7080/realms/pisciculture/protocol/openid-connect/token"

	req, err := http.PostForm(url, form)

	if err != nil {
		return nil, err
	}

	if req.StatusCode != http.StatusOK {
		return nil, errors.New("no autenticate")
	}

	var keycloak vo.KeycloakTokenVO
	if err := json.NewDecoder(req.Body).Decode(&keycloak); err != nil {
		return nil, err
	}

	return &map[string]interface{}{
		"token":      keycloak.AccessToken,
		"expires_in": keycloak.ExpiresIn,
	}, nil
}

func (a *AuthService) ValidatedAuthentication(token string) (bool, error) {
	fmt.Println(strings.Replace(token, "Bearer ", "", 1))
	var form = url.Values{
		"client_id":     {"pisciculsoft"},
		"client_secret": {"YJKvL0QvS7NL7h63Cjf5T9Y4v2tcdw39"},
		"token":         {strings.Replace(token, "Bearer ", "", 1)},
	}
	var url = "http://keycloak:7080/realms/pisciculture/protocol/openid-connect/token/introspect"

	req, err := http.PostForm(url, form)

	if err != nil {
		return false, err
	}

	if req.StatusCode != http.StatusOK {
		return false, errors.New("no autenticate")
	}

	var keycloak vo.KeycloakValidVO
	if err := json.NewDecoder(req.Body).Decode(&keycloak); err != nil {
		return false, err
	}

	return keycloak.Active, nil
}

func NewAuthService() *AuthService {
	return &AuthService{}
}
