package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/rs/cors"
)

type HttpServer struct {
	db *sql.DB
}

func (h *HttpServer) Start(port int) {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/v1/login", h.PostAuth)
	mux.HandleFunc("POST /auth/v1/register", h.Register)
	mux.HandleFunc("POST /auth/v1/mail-available", h.MailAvailable)

	mux.HandleFunc("PUT /auth/v1/mail-confirm", h.MailConfirm)

	http.ListenAndServe(":"+strconv.Itoa(port), http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			cors.AllowAll().Handler(mux).ServeHTTP(w, r)
		}))
}

func NewAuthHandlers(db *sql.DB) *HttpServer {
	return &HttpServer{
		db: db,
	}
}
