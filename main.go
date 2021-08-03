package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"time"
)

func main() {
	l := log.New(os.Stdout, "", log.LstdFlags)
	r := mux.NewRouter()
	s := http.Server{
		Addr:         ":9000",
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	err := s.ListenAndServe()
	if err != nil {
		l.Fatal(err)
	}

}
