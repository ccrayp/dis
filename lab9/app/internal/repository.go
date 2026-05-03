package internal

import (
	"app/database"
	"app/internal/logger"
	"context"
	"fmt"

	"gorm.io/gorm"
)

// Класс шаблонный репозиторий
type Repository[T any] struct {
	db        *database.Database
	logger    *logger.Logger
	tableName string
}

// Функция создания экземпляра шаблонного репозитория
func NewRepository[T any](tableName string, db *database.Database, logger *logger.Logger) IRepository[T] {
	return &Repository[T]{
		db:        db,
		logger:    logger,
		tableName: tableName,
	}
}

// Метод получения общего количества записей
func (r *Repository[T]) GetAmountQuantity() (int, error) {
	items, err := gorm.G[T](r.db.Connection).Find(context.Background())
	if err != nil {
		r.logger.MakeLog("repository", "GET", r.tableName, err.Error(), logger.Error)
		return 0, err
	}

	r.logger.MakeLog("repository", "GET", r.tableName, "get amount quantity", logger.Info)
	return len(items), nil
}

// Метод получения записей с лимитом и сдвигом (по страницам)
func (r *Repository[T]) GetByLimitOffset(limit, offset int) ([]T, error) {
	items, err := gorm.G[T](r.db.Connection).
		Limit(limit).
		Offset(offset).
		Order("id").
		Find(context.Background())
	if err != nil {
		r.logger.MakeLog("repository", "GET", r.tableName, err.Error(), logger.Error)
		return nil, err
	}

	r.logger.MakeLog("repository", "GET", r.tableName, "get by limit and offset", logger.Info)
	return items, nil
}

// Метод получения записи по идентификатору
func (r *Repository[T]) GetById(id uint) (*T, error) {
	item, err := gorm.G[T](r.db.Connection).
		Where("id = ?", id).
		First(context.Background())
	if err != nil {
		r.logger.MakeLog("repository", "GET", r.tableName, err.Error(), logger.Error)
		return nil, err
	}

	r.logger.MakeLog("repository", "GET", r.tableName, "get by id", logger.Info)
	return &item, nil
}

// Метод создания записи
func (r *Repository[T]) Create(item T) (uint, error) {
	if err := gorm.G[T](r.db.Connection).Create(context.Background(), &item); err != nil {
		r.logger.MakeLog("repository", "CREATE", r.tableName, err.Error(), logger.Error)
		return 0, err
	}

	r.logger.MakeLog("repository", "CREATE", r.tableName, fmt.Sprintf("create record. id: %d", any(item).(interface{ GetId() uint }).GetId()), logger.Info)
	return any(item).(interface{ GetId() uint }).GetId(), nil
}

// Метод обновления данных
func (r *Repository[T]) Update(item *T) error {
	_, err := gorm.G[T](r.db.Connection).
		Where("id = ?", any(item).(interface{ GetId() uint }).GetId()).
		Updates(context.Background(), *item)

	if err != nil {
		r.logger.MakeLog("repository", "UPDATE", r.tableName, fmt.Sprintf("id: %d. errod: %s", any(item).(interface{ GetId() uint }).GetId(), err.Error()), logger.Error)
	}

	r.logger.MakeLog("repository", "UPDATE", r.tableName, fmt.Sprintf("update record. id: %d", any(item).(interface{ GetId() uint }).GetId()), logger.Info)
	return err
}

// Метод удаления записи
func (r *Repository[T]) Delete(id uint) error {
	_, err := gorm.G[T](r.db.Connection).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		r.logger.MakeLog("repository", "DELETE", r.tableName, fmt.Sprintf("id: %d. error: %s", id, err.Error()), logger.Error)
	}

	r.logger.MakeLog("repository", "DELETE", r.tableName, fmt.Sprintf("id: %d", id), logger.Info)
	return err
}
