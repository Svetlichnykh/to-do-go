package Ucmds

import (
	"fmt"
	"to-do/logs"
)

func Exit(text []string) bool {
	if len(text) > 1 {
		logs.NewLog(logs.Counter, "Пользователь ввел exit с лишними аргументами")
		fmt.Println("Команда exit не принимает аргументов, введите help для вывода доступных команд")
		return false
	}
	logs.NewLog(logs.Counter, "Пользователь завершил выполнение программы")
	fmt.Println("ВЫХОД...")
	return true
}
