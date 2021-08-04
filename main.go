package main

import (
	"context"
	"fmt"
	"github.com/vkParser/models"
	"github.com/vkParser/rabbitmq"
	"github.com/vkParser/server"
	"log"
	"net/http"
	"os"
	"os/signal"
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

	go func (){
		err := rabbitmq.GetTask()
		if err!=nil{
			fmt.Println(err)
		}
	}()

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	sig := <-sigChan
	l.Println("GraceFull shutdown ", sig)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
