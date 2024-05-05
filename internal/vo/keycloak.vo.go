package vo

type KeycloakValidVO struct {
	Active bool `json:"active"`
}

type KeycloakTokenVO struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
