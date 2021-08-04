package main

import (
	"github.com/vkParser/models"
	"github.com/vkParser/server"
	"log"
	"net/http"
	"os"
	"time"
)



func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)
	env := models.NewEnv(l)
	r := server.NewRouter(env)

	s := http.Server{
		Addr:         ":9000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

}
