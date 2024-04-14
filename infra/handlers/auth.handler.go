package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/server/auth/internal/core/service"
	"github.com/server/auth/internal/dao"
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

	tx, err := h.db.Begin()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authDao := dao.NewAuthDao(h.db)
	loginDao := dao.NewLoginDao(tx)
	userService := service.NewAuthService(authDao, loginDao)

	userAuth, err := userService.Auth(auth.Mail, auth.Password)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(userAuth)
}
