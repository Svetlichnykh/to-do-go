package tasks

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"to-do/logs"
)

func HandleChange(text []string) {

	var title string

	var newTitle string
	var newDescription string
	var newCategory string
	var newTargetDate string
	var newCreationDate string
	var newDoneDate string

	var errInput error

	flag := "title"

	for i, v := range text {
		if i != 0 {
			switch v {
			case "-nt":
				if newTitle != "" {
					errInput = errors.Join(errInput, errors.New("флаг -nt был введен более 1 раза"))
				} else {
					flag = "newTitle"
				}
			case "-nd":
				if newDescription != "" {
					errInput = errors.Join(errInput, errors.New("флаг -nd был введен более 1 раза"))
				} else {
					flag = "newDescription"
				}
			case "-nc":
				if newCategory != "" {
					errInput = errors.Join(errInput, errors.New("флаг -nc был введен более 1 раза"))
				} else {
					flag = "newCategory"
				}
			case "-ntd":
				if newTargetDate != "" {
					errInput = errors.Join(errInput, errors.New("флаг -ntd был введен более 1 раза"))
				} else {
					flag = "newTargetDate"
				}
			case "-ncd":
				if newCreationDate != "" {
					errInput = errors.Join(errInput, errors.New("флаг -ncd был введен более 1 раза"))
				} else {
					flag = "newCreationDate"
				}
			case "-ndd":
				if newDoneDate != "" {
					errInput = errors.Join(errInput, errors.New("флаг -ndt был введен более 1 раза"))
				} else {
					flag = "newDoneDate"
				}
			default:
				switch flag {
				case "title":
					title += v + " "
				case "newTitle":
					newTitle += v + " "
				case "newDescription":
					newDescription += v + " "
				case "newCategory":
					newCategory += v + " "
				case "newTargetDate":
					newTargetDate += v + " "
				case "newCreationDate":
					newCreationDate += v + " "
				case "newDoneDate":
					newDoneDate += v + " "
				}
			}
		}
	}

	if errInput != nil {
		fmt.Println(errInput)
		return
	}

	var ntDateTime time.Time
	if newTargetDate != "" {
		ntDateTime = TimeTranslate(newTargetDate)
		if ntDateTime.IsZero() {
			logText := "Пользователь допустил ошибку при вводе времени дедлайна во время изменения задачи - " + title
			logs.NewLog(logs.Counter, logText)
			fmt.Println("Новая дата дедлайна введена неверно, формат - 2006.01.02 15:04")
			return
		}
	}

	var ncDateTime time.Time
	if newCreationDate != "" {
		ncDateTime = TimeTranslate(newCreationDate)
		if ncDateTime.IsZero() {
			logText := "Пользователь допустил ошибку при вводе времени создания во время изменения задачи - " + title
			logs.NewLog(logs.Counter, logText)
			fmt.Println("Новая дата создания введена неверно, формат - 2006.01.02 15:04")
			return
		}
	}

	var ndDateTime time.Time
	if newDoneDate != "" {
		ndDateTime = TimeTranslate(newDoneDate)
		if ndDateTime.IsZero() {
			logText := "Пользователь допустил ошибку при вводе времени выполнения во время изменения задачи - " + title
			logs.NewLog(logs.Counter, logText)
			fmt.Println("Новая дата выполнения введена неверно, формат - 2006.01.02 15:04")
			return
		}
	}

	newOptions := Task{
		title:        newTitle,
		description:  newDescription,
		category:     newCategory,
		done:         false,
		creationDate: ncDateTime,
		targetDate:   ntDateTime,
		doneDate:     ndDateTime,
	}

	for _, v := range Pool {
		if strings.ToLower(v.title) == strings.ToLower(newTitle) {
			logs.NewLog(logs.Counter, "Пользователь попытался изменить название задачи на уже существующее")
			fmt.Println("Задача с таким названием уже существует! Введите другое новое имя")
			return
		}
	}

	if title == "" {
		logs.NewLog(logs.Counter, "Пользователь попытался изменить задачу, но не ввел ее название")
		fmt.Println("Вы не ввели название! Введите help для вывода списка доступных команд")
		return
	} else {
		EditTask(newOptions, title)
	}
}

func TimeTranslate(strTime string) time.Time {

	if strTime == "" {
		return time.Time{}
	}

	typeTime, err := time.Parse("2006.01.02 15:04", strings.TrimSpace(strTime))

	if err != nil {
		return time.Time{}
	}

	return typeTime
}
