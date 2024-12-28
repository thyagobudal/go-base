package server

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{logger}
}
