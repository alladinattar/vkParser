package server

import (
	"github.com/gorilla/mux"
	"github.com/vkParser/handlers"
	"github.com/vkParser/models"
)

func NewRouter(e *models.Env) *mux.Router {
	grabberHandler := handlers.NewGrabberHandler(e)
	r := mux.NewRouter()
	r.Handle("/api/v1/graber_profiles", grabberHandler).Methods("GET")
	r.Handle("/api/v1/config", nil).Methods("PUT")
	r.Handle("/api/v1/configs",nil).Methods("GET")
	r.Handle("/api/v1/accounts_add", nil).Methods("POST")
	r.Handle("/api/v1/accounts_checker", nil).Methods("GET")
	r.Handle("/api/v1/accounts", nil).Methods("GET")
	r.Handle("/api/v1/vk_auth_profile", nil).Methods("POST")
	r.Handle("/api/v1/profile_notifications", nil).Methods("POST")
	r.Handle("/api/v1/profile_likes", nil).Methods("GET")
	r.Handle("/api/v1/profile_comments", nil).Methods("GET")
	return r
}
