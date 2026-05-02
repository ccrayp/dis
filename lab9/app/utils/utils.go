package utils

import "fmt"

func Pause() {
	fmt.Print("\n\nНажмите Enter...")
	fmt.Scanln()
}

func Clean() {
	fmt.Print("\033[2J\033[H")
}
