package handlers

import (
	"encoding/json"
	"github.com/vkParser/server"
	"log"
	"net/http"
)

type GrabberHandler struct {
	l *log.Logger
}

func NewGrabberHandler(env *server.Env) *GrabberHandler {
	return &GrabberHandler{l: env.L}
}

func (b *GrabberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct{
		id string
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err!=nil{
		b.l.Println("cannot decode vk_id")
	}

}
