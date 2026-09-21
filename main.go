package main

import (
	"fmt"
	"strings"
	"to-do/Ucmds"
	"to-do/input"
	"to-do/logs"
	"to-do/tasks"
)

var Text []string

func main() {
	logs.NewLog(logs.Counter, "Пользователь запустил программу")
	logs.CounterInc()

	fmt.Println("========== HELLO USER ==========")
	fmt.Println("==== WELCOME TO MY TODO APP ====")

	for {
		fmt.Print("> ")

		if ok := input.Scanner.Scan(); !ok {
			logs.NewLog(logs.Counter, "Ошибка ввода. Прекращение работы программы")
			fmt.Println("Ошибка")
			return
		}

		Text = strings.Fields(input.Scanner.Text())

		if len(Text) == 0 {

			logs.NewLog(logs.Counter, "Пользователь ввел пустую строку")
			Ucmds.EmtyInput()

		} else {
			cmd := Text[0]

			switch cmd {
			// UTILS
			case "exit":
				if exitHandler := Ucmds.Exit(Text); exitHandler {
					return
				}
			case "help":
				Ucmds.Help(Text)
			case "logs":
				Ucmds.ShowLogs(Text)
			// TASKS
			case "add":
				tasks.HandleAdd(Text)
			case "list":
				tasks.HandleList(Text)
			case "del":
				logs.NewLog(logs.Counter, "Пользователь удалил задачу")
			case "done":
				logs.NewLog(logs.Counter, "Пользователь отметил задачу как выполненную")
			// WRONG
			default:
				Ucmds.WrongInput(Text)
			}

		}

		logs.CounterInc()
	}
}
