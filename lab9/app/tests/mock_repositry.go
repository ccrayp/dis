package tests

import (
	"fmt"
)

// Моковый шаблонный репозиторий с ошибками
type ErrorRepository[T any] struct{}

func (r *ErrorRepository[T]) GetAmountQuantity() (int, error) {
	return 0, fmt.Errorf("mock error")
}

func (r *ErrorRepository[T]) GetByLimitOffset(limit, offset int) ([]T, error) {
	return nil, fmt.Errorf("mock error")
}

func (r *ErrorRepository[T]) GetById(id uint) (*T, error) {
	return nil, fmt.Errorf("mock error")
}

func (r *ErrorRepository[T]) Create(item T) (uint, error) {
	return 0, fmt.Errorf("mock error")
}

func (r *ErrorRepository[T]) Update(item *T) error {
	return fmt.Errorf("mock error")
}

func (r *ErrorRepository[T]) Delete(id uint) error {
	return fmt.Errorf("mock error")
}
