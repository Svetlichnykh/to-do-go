package tasks

import (
	"fmt"
	"to-do/logs"

	"github.com/k0kubun/pp"
)

func HandleList(text []string) {
	if len(text) > 1 {
		logs.NewLog(logs.Counter, "Пользователь ввел list с лишними аргументами")
		fmt.Println("Команда list не принимает аргументов, введите help для вывода доступных команд")
	} else {
		logs.NewLog(logs.Counter, "Пользователь вывел список заданий")
		pp.Println(Pool)
	}
}
