package server

import (
	"github.com/gorilla/mux"
	"github.com/vkParser/handlers"
	"log"
)

func NewRouter(e *Env)*mux.Router{
	grabberHandler := handlers.NewGrabberHandler(e)
	r := mux.NewRouter()
	r.Handle("/api/v1/graber_profiles", grabberHandler)

	return r
}


type Env struct {
	L *log.Logger
}

func NewEnv(l *log.Logger) *Env {
	return &Env{l}
}
