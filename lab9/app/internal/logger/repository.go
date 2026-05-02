package logger

import (
	"app/database"
	"app/model"
	"context"

	"gorm.io/gorm"
)

// Класс репозиторий для доступа к таблице audit в БД
type AuditRepository struct {
	db *database.Database
}

// Функция создания экземпляра репозитория журанла аудита
func NewAuditRepository(db *database.Database) *AuditRepository {
	return &AuditRepository{
		db: db,
	}
}

// Метод созадния записи в журнал
func (r *AuditRepository) CreateRecord(record *model.Audit) error {
	return gorm.G[model.Audit](r.db.Connection).Create(context.Background(), record)
}
