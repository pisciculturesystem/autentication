package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/server/auth/internal/core/service"
	"github.com/server/auth/internal/dao"
	"github.com/server/auth/internal/vo"
)

func (h *HttpServer) Register(w http.ResponseWriter, r *http.Request) {
	var data vo.UserCreatedVO

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	userDao := dao.NewUserDao(h.db)
	mailDao := dao.NewMailDao(h.db)
	userService := service.NewUserService(userDao, mailDao)

	id, err := userService.Create(&data)

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
