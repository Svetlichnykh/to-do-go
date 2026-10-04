package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"to-do/cmds"
	"to-do/logs"
	"to-do/tasks"
)

var Text []string
var Scanner = bufio.NewScanner(os.Stdin)

func main() {
	logs.NewLog(logs.Counter, "Пользователь запустил программу")
	logs.CounterInc()

	fmt.Println("========== HELLO USER ==========")
	fmt.Println("==== WELCOME TO MY TODO APP ====")

	for {
		fmt.Print("> ")

		if ok := Scanner.Scan(); !ok {
			logs.NewLog(logs.Counter, "Ошибка ввода. Прекращение работы программы")
			fmt.Println("Ошибка")
			return
		}

		Text = strings.Fields(Scanner.Text())

		if len(Text) == 0 {

			logs.NewLog(logs.Counter, "Пользователь ввел пустую строку")
			cmds.EmtyInput()

		} else {
			cmd := Text[0]

			switch cmd {
			// UTILS
			case "exit":
				if exitHandler := cmds.Exit(Text); exitHandler {
					return
				}
			case "help":
				cmds.Help(Text)
			case "logs":
				cmds.ShowLogs(Text)
			// TASKS
			case "add":
				tasks.HandleAdd(Text)
			case "list":
				tasks.HandleList(Text)
			case "del":
				tasks.HandleDel(Text)
			case "done":
				tasks.HandleDone(Text)
			case "undone":
				tasks.HandleUndone(Text)
			case "change":
				tasks.HandleChange(Text)
			// WRONG
			default:
				cmds.WrongInput(Text)
			}

		}

		logs.CounterInc()
	}
}
