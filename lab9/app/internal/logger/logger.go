package logger

import (
	"app/model"
	"time"
)

type Level string

const (
	Info    Level = "info"
	Warning Level = "warning"
	Error   Level = "error"
)

type Logger struct {
	repository AuditRepository
}

func NewLogger(repository AuditRepository) *Logger {
	return &Logger{
		repository: repository,
	}
}

func (l *Logger) MakeLog(actor, action, target, descirption string, level Level) {
	record := model.Audit{
		Actor:       actor,
		Action:      action,
		Target:      target,
		Level:       string(level),
		Description: descirption,
		Timestamp:   time.Now(),
	}

	l.repository.CreateRecord(&record)
}
