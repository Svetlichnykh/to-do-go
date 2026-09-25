package tasks

import (
	"fmt"
	"strings"
	"time"
	"to-do/logs"

	"github.com/k0kubun/pp"
)

func HandleAdd(text []string) {

	var title string
	var description string
	var category string
	var targetDate string

	flag := "title"

	for i, v := range text {
		if i != 0 {
			switch v {
			case "-d":
				flag = "description"
			case "-c":
				flag = "category"
			case "-t":
				flag = "targetDate"
			default:
				switch flag {
				case "title":
					title += v + " "
				case "description":
					description += v + " "
				case "category":
					category += v + " "
				case "targetDate":
					targetDate += v + " "
				}
			}
		}
	}

	var tDateTime time.Time
	var err error

	if targetDate != "" {
		tDateTime, err = time.Parse(
			"15:04 02.01.2006",
			strings.TrimSpace(targetDate),
		)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	options := Options{
		Description: strings.TrimSpace(description),
		Category:    strings.TrimSpace(category),
		TargetDate:  tDateTime,
	}

	if title == "" {
		logs.NewLog(logs.Counter, "Пользователь попытался создать задачу без названия")
		fmt.Println("Вы не ввели название! Введите help для вывода списка доступных команд")
		return
	} else {
		NewTask(strings.TrimSpace(title), options)
	}
}
