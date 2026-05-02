package model

import "time"

type Audit struct {
	Id          uint
	Actor       string
	Action      string
	Target      string
	Level       string
	Description string
	Timestamp   time.Time
}

func (Audit) TableName() string {
	return "audit"
}
