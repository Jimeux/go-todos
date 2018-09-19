package common

import (
	"github.com/fluent/fluent-logger-golang/fluent"
)

type Logger interface {
	Forward(tag string, data map[string]interface{})
	Close()
}

func NewLogger(client *fluent.Fluent) Logger {
	return LoggerImpl{client}
}

type LoggerImpl struct {
	client *fluent.Fluent
}

func (l LoggerImpl) Forward(tag string, data map[string]interface{}) {
	l.client.Post(tag, data)
}

func (l LoggerImpl) Close() {
	l.client.Close()
}
