package main

import (
	"fmt"
	"strconv"
	"time"
	"to-do/input"

	"github.com/k0kubun/pp"
)

type Task struct {
	title        string
	description  string
	category     string
	done         bool
	creationDate time.Time
	doneDate     time.Time
}

func main() {
	pool := []Task{}
	logs := []string{"======== Логи программы ========"}
	counter := 1
	hollowsCount := 0
	current := strconv.Itoa(counter) + ". "
	logs = append(logs, current+"Пользователь запустил программу")
	counter++

	fmt.Println("========== HELLO USER ==========")
	fmt.Println("==== WELCOME TO MY TODO APP ====")

	for {

		current = strconv.Itoa(counter) + ". "
		fmt.Print("> ")

		if ok := input.Scanner.Scan(); !ok {
			logs = append(logs, current+"Ошибка ввода. Прекращение работы программы")
			fmt.Println("Ошибка")
			return
		}

		text := input.Scanner.Text()

		if text == "" {

			if hollowsCount >= 3 {
				logs = append(logs, current+"Пользователь ввел пустую строку 3+ раза подряд")
				pp.Println("Идиот, хватит тыкать ENTER")
			}

			logs = append(logs, current+"Пользователь ввел пустую строку")
			fmt.Println("Пожалуйста, введите команду")
			hollowsCount++

		} else {
			hollowsCount = 0
			fmt.Println("Ваша команда:", text)

			// Utils

			if text == "exit" {
				logs = append(logs, current+"Пользователь завершил выполнение программы")
				fmt.Println("ВЫХОД...")
				return
			}

			if text == "help" {
				logs = append(logs, current+"Пользователь ввел help и получил список доступных программ")
				fmt.Println("")
				fmt.Println("Список команд:")
				fmt.Println("")
				fmt.Println("help — эта команда позволяет узнать доступные команды и их формат")
				fmt.Println("")
				fmt.Println("add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
				fmt.Println("")
				fmt.Println("list — эта команда позволяет получить полный список всех задач")
				fmt.Println("")
				fmt.Println("del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
				fmt.Println("")
				fmt.Println("done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
				fmt.Println("")
				fmt.Println("log — эта команда позволяет получить лог текущего выполнения")
				fmt.Println("")
				fmt.Println("exit — эта команда позволяет завершить выполнение программы")
				fmt.Println("")
			}

			if text == "logs" {
				logs = append(logs, current+"Пользователь смотрит логи программы")
				for _, v := range logs {
					fmt.Println(v)
				}
			}

			// Tasks

			if text == "add" {
				logs = append(logs, current+"Пользователь добавляет задачу")
				fmt.Println("Добавляю задачу...")
			}

			if text == "list" {
				logs = append(logs, current+"Пользователь вывел список всех задач")
				fmt.Println("Список задач:")
				pp.Println(pool)
			}

		}

		counter++

	}
}
