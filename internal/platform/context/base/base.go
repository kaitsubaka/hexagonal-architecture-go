package base

import (
	"context"
	"os"
	"os/signal"
)

type key int

const (
	CtxEnvKey key = iota
)

func New() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.WithValue(context.Background(), CtxEnvKey, os.Getenv("APP_ENV")), os.Interrupt)
}
