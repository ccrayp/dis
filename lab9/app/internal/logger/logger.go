package logger

import (
	"app/model"
	"time"
)

// Уровень записи в журнале
type Level string

// Перечисление уровней записи в журале
const (
	Info    Level = "info"
	Warning Level = "warning"
	Error   Level = "error"
)

// Класс логгер
type Logger struct {
	repository AuditRepository
}

// Функция создания экземпляра логгера
func NewLogger(repository AuditRepository) *Logger {
	return &Logger{
		repository: repository,
	}
}

// Метод записи события в журнал аудита
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
