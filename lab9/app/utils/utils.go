package utils

import "fmt"

// Функция приоставновки выполнения программы до нажатия любой кнопки
func Pause() {
	fmt.Print("\n\nНажмите Enter...")
	fmt.Scanln()
}

// Функция очистки экрана
func Clean() {
	fmt.Print("\033[2J\033[H")
}
