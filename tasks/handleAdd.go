package tasks

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"to-do/logs"
)

func HandleAdd(text []string) {

	var title string
	var description string
	var category string
	var targetDate string
	var errInput error

	flag := "title"

	for i, v := range text {
		if i != 0 {
			switch v {
			case "-d":
				if description != "" {
					errInput = errors.Join(errInput, errors.New("флаг -d был введен более 1 раза"))
				} else {
					flag = "description"
				}
			case "-c":
				if category != "" {
					errInput = errors.Join(errInput, errors.New("флаг -c был введен более 1 раза"))
				} else {
					flag = "category"
				}
			case "-t":
				if targetDate != "" {
					errInput = errors.Join(errInput, errors.New("флаг -t был введен более 1 раза"))
				} else {
					flag = "targetDate"
				}
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

	if errInput != nil {
		fmt.Println(errInput)
		return
	}

	var tDateTime time.Time
	var err error

	if targetDate != "" {
		tDateTime, err = time.Parse(
			"2006.01.02 15:04",
			strings.TrimSpace(targetDate),
		)
	}

	if err != nil {
		logs.NewLog(logs.Counter, "Пользователь допустил ошибку при вводе времени дедлайна во время создания новой задачи")
		fmt.Println("Дата введена неверно, формат - 2006.01.02 15:04")
		return
	}

	options := Options{
		Description: strings.TrimSpace(description),
		Category:    strings.TrimSpace(category),
		TargetDate:  tDateTime,
	}
	title = strings.TrimSpace(title)

	for _, v := range Pool {
		if strings.ToLower(v.title) == strings.ToLower(title) {
			logs.NewLog(logs.Counter, "Пользователь попытался создать уже существующую задачу")
			fmt.Println("Задача с таким названием уже существует! Введите другое")
			return
		}
	}

	if title == "" {
		logs.NewLog(logs.Counter, "Пользователь попытался создать задачу без названия")
		fmt.Println("Вы не ввели название! Введите help для вывода списка доступных команд")
		return
	} else {
		NewTask(title, options)
	}
}
