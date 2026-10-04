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

var text []string

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	logs.NewLog(logs.Counter, "Пользователь запустил программу")
	logs.CounterInc()

	fmt.Println("========== HELLO USER ==========")
	fmt.Println("==== WELCOME TO MY TODO APP ====")

	for {
		fmt.Print("> ")

		if ok := scanner.Scan(); !ok {
			logs.NewLog(logs.Counter, "Ошибка ввода. Прекращение работы программы")
			fmt.Println("Ошибка")
			return
		}

		text = strings.Fields(scanner.Text())

		if len(text) == 0 {
			logs.NewLog(logs.Counter, "Пользователь ввел пустую строку")
			cmds.EmptyInput()
			continue
		}

		cmd := text[0]

		switch cmd {
		// UTILS
		case "exit":
			if cmds.Exit(text) {
				return
			}
		case "help":
			cmds.Help(text)
		case "logs":
			cmds.ShowLogs(text)
		// TASKS
		case "add":
			tasks.HandleAdd(text)
		case "list":
			tasks.HandleList(text)
		case "del":
			tasks.HandleDel(text)
		case "done":
			tasks.HandleDone(text)
		case "undone":
			tasks.HandleUndone(text)
		case "change":
			tasks.HandleChange(text)
		// WRONG
		default:
			cmds.WrongInput(text)
		}

		logs.CounterInc()
	}
}
