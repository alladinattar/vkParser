package server

import (
	"github.com/gorilla/mux"
	"github.com/vkParser/handlers"
	"github.com/vkParser/models"
)

func NewRouter(e *models.Env) *mux.Router {
	grabberHandler := handlers.NewGrabberHandler(e)
	r := mux.NewRouter()
	r.Handle("/api/v1/graber_profiles", grabberHandler)

	return r
}
