package models

import "log"

type Env struct {
	L *log.Logger
}

func NewEnv(l *log.Logger) *Env {
	return &Env{l}
}
