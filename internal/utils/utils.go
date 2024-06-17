package utils

import (
	"log"
	"os"
)

// CheckErr обрабатывает ошибки и завершает выполнение программы при необходимости
func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

// FileExists проверяет, существует ли файл по указанному пути
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
