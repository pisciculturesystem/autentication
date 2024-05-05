package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/server/auth/internal/core/service"
	"github.com/server/auth/internal/vo"
)

func (h *HttpServer) PostAuth(w http.ResponseWriter, r *http.Request) {

	var auth vo.LoginVO
	if err := json.NewDecoder(r.Body).Decode(&auth); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	if err := auth.IsValid(); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	session, err := service.NewAuthService().Auth(auth.Username, auth.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(session)
}

func (h *HttpServer) ValidatedAuthentication(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	valid, err := service.NewAuthService().ValidatedAuthentication(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if valid {
		w.WriteHeader(http.StatusAccepted)
		return
	}
	w.WriteHeader(http.StatusForbidden)
}
