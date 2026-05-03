package internal

// Интерфейс шаблонного репозитория
type IRepository[T any] interface {
	GetAmountQuantity() (int, error)
	GetByLimitOffset(limit, offset int) ([]T, error)
	GetById(id uint) (*T, error)
	Create(item T) (uint, error)
	Update(item *T) error
	Delete(id uint) error
}
