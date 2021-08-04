package handlers

import (
	"encoding/json"
	"github.com/vkParser/models"
	"github.com/vkParser/rabbitmq"
	"log"
	"net/http"
)

type GrabberHandler struct {
	l *log.Logger
}

func NewGrabberHandler(env *models.Env) *GrabberHandler {
	return &GrabberHandler{l: env.L}
}

func (b *GrabberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Id string `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		b.l.Println("cannot decode vk_id")
	}
	err = rabbitmq.AddTask([]byte(body.Id))
	if err!=nil{
		b.l.Println(err)
	}
}
