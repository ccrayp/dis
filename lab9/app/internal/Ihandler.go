package internal

// Интерфейс обработчика для меню
type Handler interface {
	Menu()
	Name() string
}
