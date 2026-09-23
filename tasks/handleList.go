package tasks

import (
	"fmt"
	"strconv"
	"strings"
	"to-do/logs"
	"to-do/timeCalc"
)

func HandleList(text []string) {
	if len(text) > 1 {
		if text[1] == "-c" {
			var cat string
			for i, v := range text {
				if i >= 2 {
					cat += v + " "
				}
			}
			cat = strings.TrimSpace(cat)

			fmt.Println("======== СПИСОК ЗАДАЧ ========")
			fmt.Println("Поиск по категории:", cat)

			counter := 0
			for _, v := range Pool {
				if strings.ToLower(v.category) == strings.ToLower(cat) {
					counter++
				}
			}

			if counter == 0 {
				logText := "Пользователь вывел список задач по категории, но их не оказалось - " + cat
				logs.NewLog(logs.Counter, logText)
				fmt.Println("ЗАДАЧИ В ЭТОЙ КАТЕГОРИИ НЕ НАЙДЕНЫ")
				return
			} else {
				logText := "Пользователь вывел список задач по категории, но их не оказалось - " + cat
				logs.NewLog(logs.Counter, logText)
				fmt.Println("Всего задач:", strconv.Itoa(counter))
			}

			fmt.Println("")

			i := 0
			for _, v := range Pool {
				if strings.ToLower(v.category) == strings.ToLower(cat) {
					PrintTask(i, v)
					i++
				}
			}

		} else {
			logs.NewLog(logs.Counter, "Пользователь ввел list с лишними аргументами")
			fmt.Println("Команда list не принимает аргументов, введите help для вывода доступных команд")
		}

	} else {

		fmt.Println("======== СПИСОК ЗАДАЧ ========")
		fmt.Println("")
		if len(Pool) == 0 {
			fmt.Println("ЗАДАЧ НЕТ")
			logs.NewLog(logs.Counter, "Пользователь вывел список задач, но задач нету")
		} else {
			fmt.Println("Всего задач:", strconv.Itoa(len(Pool)))
			logs.NewLog(logs.Counter, "Пользователь вывел список задач")
		}

		fmt.Println("")
		for i, v := range Pool {
			PrintTask(i, v)
		}
	}
}

func PrintTask(i int, v Task) {
	fmt.Print(strconv.Itoa(i+1) + ". ")
	if v.category != "" {
		fmt.Print("(", v.category, ") ")
	}
	fmt.Print("Задача - ", v.title)
	var check string
	if v.done {
		check = " [ ✅ ]"
	} else {
		check = " [ ❌ ]"
	}
	fmt.Println(check)

	if v.description != "" {
		fmt.Println(v.description)
	}

	fmt.Println("📅 Дата создания:", v.creationDate.Format("2006.01.02 15:04"))
	if !v.targetDate.IsZero() {
		fmt.Print("❗ Дедлайн:", v.targetDate.Format("2006.01.02 15:04"))
		fmt.Println(" Осталось времени: " + timeCalc.TimeUntil(v.targetDate))
	}
	if !v.doneDate.IsZero() {
		fmt.Println("✅ Время выполнения:", v.doneDate.Format("2006.01.02 15:04"))
	}
	fmt.Println("")
}
