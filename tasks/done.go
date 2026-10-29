package tasks

import (
	"fmt"
	"time"
	"to-do/logs"
)

func HandleDone(text []string) {
	title := titleFrom(text)

	i, err := findTask(title)

	if err != nil {
		fmt.Println(err)
		logs.NewLog(logs.Counter, "Пользователь попытался отметить задачу как выполненную, но получил ошибку:")
		logs.NewLog(0, err.Error())
		return
	}

	if Pool[i].done {
		fmt.Println("Данная задача уже отмечена как выполненная - " + title)
		logText := "Пользователь попытался отметить уже выполненную задачу выполненной - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}

	Pool[i].doneDate = time.Now()
	Pool[i].done = true
	fmt.Println("Задача выполнена - " + title)
	logText := "Пользователь отметил следующую задачу как выполненную - " + title
	logs.NewLog(logs.Counter, logText)

}
