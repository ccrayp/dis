package logger

import (
	"app/database"
	"app/model"
	"context"

	"gorm.io/gorm"
)

type AuditRepository struct {
	db *database.Database
}

func NewAuditRepository(db *database.Database) *AuditRepository {
	return &AuditRepository{
		db: db,
	}
}

func (r *AuditRepository) CreateRecord(record *model.Audit) error {
	return gorm.G[model.Audit](r.db.Connection).Create(context.Background(), record)
}
