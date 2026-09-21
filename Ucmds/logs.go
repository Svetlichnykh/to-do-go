package Ucmds

import (
	"fmt"
	"to-do/logs"
)

func ShowLogs(text []string) {
	if len(text) > 1 {
		logs.NewLog(logs.Counter, "Пользователь ввел logs с лишними аргументами")
		fmt.Println("Команда logs не принимает аргументов, введите help для вывода доступных команд")
	} else {
		logs.NewLog(logs.Counter, "Пользователь посмотрел логи программы")
		fmt.Println("Ваша команда - logs")
		fmt.Println("======== Логи программы ========")
		fmt.Println("")
		for _, v := range logs.Logs {
			fmt.Println(v)
		}
	}
}
