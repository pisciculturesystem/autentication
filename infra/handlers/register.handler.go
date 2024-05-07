package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/server/auth/internal/core/service"
	"github.com/server/auth/internal/dao"
	"github.com/server/auth/internal/vo"
)

func (h *HttpServer) Register(w http.ResponseWriter, r *http.Request) {

	var userVo vo.KeycloakUserVO
	if err := json.NewDecoder(r.Body).Decode(&userVo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	id, err := service.NewUserService().Create(&userVo)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

func (h *HttpServer) MailAvailable(w http.ResponseWriter, r *http.Request) {

	var mail vo.MailAvailableVO

	if err := json.NewDecoder(r.Body).Decode(&mail); err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	mailDao := dao.NewMailDao(h.db)
	mailService := service.NewMailService(mailDao)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"available": !mailService.IsAvailable(mail.Mail),
	})
}

func (h *HttpServer) MailConfirm(w http.ResponseWriter, r *http.Request) {

}
