package util

import (
	"log"
)

const (
	COLOR_RESET  = "\033[0m"
	COLOR_RED    = "\033[31m"
	COLOR_GREEN  = "\033[32m"
	COLOR_YELLOW = "\033[33m"
	COLOR_BLUE   = "\033[34m"
	COLOR_PURPLE = "\033[35m"
	COLOR_CYAN   = "\033[36m"
	COLOR_GRAY   = "\033[37m"
	COLOR_WHITE  = "\033[97m"
)

func PrintRedError(a ...any) {
	log.Println(COLOR_RED, a, COLOR_RESET)
}

func PrintGreenStatus(a ...any) {
	log.Println(COLOR_GREEN, a, COLOR_RESET)
}

func PrintYellowStatus(a ...any) {
	log.Println(COLOR_YELLOW, a, COLOR_RESET)
}

func PrintPurpleWarning(a ...any) {
	log.Println(COLOR_PURPLE, a, COLOR_RESET)
}
