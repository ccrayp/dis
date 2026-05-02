package model

import "time"

// Структура описывающая модель данных журанал
type Audit struct {
	Id          uint
	Actor       string
	Action      string
	Target      string
	Level       string
	Description string
	Timestamp   time.Time
}

// Метод получения имени таблицы журанла
func (Audit) TableName() string {
	return "audit"
}
