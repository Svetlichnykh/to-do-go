package tasks

import (
	"fmt"
	"time"
	"to-do/logs"
)

func HandleUndone(text []string) {
	title := titleFrom(text)

	i, err := findTask(title)

	if err != nil {
		fmt.Println(err)
		logs.NewLog(logs.Counter, "Пользователь попытался отметить задачу как невыполненную, но получил ошибку:")
		logs.NewLog(0, err.Error())
		return
	}

	if !Pool[i].done {
		fmt.Println("Данная задача уже отмечена как невыполненная - " + title)
		logText := "Пользователь попытался отметить невыполненную задачу невыполненной - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}

	Pool[i].doneDate = time.Time{}
	Pool[i].done = false
	fmt.Println("Задача отмечена как не выполненная - " + title)
	logText := "Пользователь отметил следующую задачу как невыполненную - " + title
	logs.NewLog(logs.Counter, logText)
}
