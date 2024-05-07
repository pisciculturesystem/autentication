package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/server/auth/internal/vo"
)

type UserService struct {
}

func (a *UserService) Create(input *vo.KeycloakUserVO) (string, error) {
	if err := input.IsValid(); err != nil {
		return "", err
	}

	var urlAuth = "http://keycloak:7080/realms/pisciculture/protocol/openid-connect/token"
	var urlCreateUser = "http://keycloak:7080/auth/admin/realms/pisciculture/users"

	var auth = url.Values{
		"client_id":     {"pisciculsoft"},
		"client_secret": {"YJKvL0QvS7NL7h63Cjf5T9Y4v2tcdw39"},
		"grant_type":    {"password"},
		"username":      {"dev"},
		"password":      {"123"},
	}

	req, err := http.PostForm(urlAuth, auth)
	if err != nil {
		return "", err
	}

	if req.StatusCode != http.StatusOK {
		return "", errors.New("no autenticate")
	}

	var keycloak vo.KeycloakTokenVO
	if err := json.NewDecoder(req.Body).Decode(&keycloak); err != nil {
		return "", err
	}

	userJSON, err := json.Marshal(map[string]interface{}{
		"username":      input.Username,
		"enabled":       true,
		"emailVerified": true,
		"firstName":     input.FirstName,
		"lastName":      input.LastName,
		"email":         input.Mail,
		"credentials": []interface{}{
			map[string]interface{}{
				"type":      "password",
				"value":     input.Password,
				"temporary": false,
			},
		},
	})

	rq, err := http.NewRequest("POST", urlCreateUser, bytes.NewBuffer(userJSON))
	if err != nil {
		return "", err
	}

	rq.Header.Set("Content-Type", "application/json")
	rq.Header.Set("Authorization", "Bearer "+keycloak.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(rq)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return "", err
	}

	return "", nil
}

func NewUserService() *UserService {
	return &UserService{}
}
