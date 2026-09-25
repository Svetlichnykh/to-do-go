package tasks

import (
	"fmt"
	"strings"
	"time"
	"to-do/logs"
)

func HandleUndone(text []string) {
	var title string
	for i := range text {
		if i > 0 {
			title += text[i] + " "
		}
	}
	title = strings.TrimSpace(title)
	var logText string
	existFlag := false

	for i := range Pool {
		v := &Pool[i]
		if strings.ToLower(v.title) == strings.ToLower(title) {
			existFlag = true
			if v.done == false {
				fmt.Println("Данная задача уже отмечена как невыполненная - " + title)
				logText = "Пользователь попытался отметить невыполненную задачу невыполненной - " + title
				logs.NewLog(logs.Counter, logText)
				return
			}
			v.doneDate = time.Time{}
			v.done = false
			logText = "Пользователь отметил следующую задачу как невыполненную - " + title
		}
	}

	if !existFlag {
		fmt.Println("Не найдено задачи с названием", title)
		logText = "Пользователь попытался отметить несуществующую задачу как невыполненную - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}

	logs.NewLog(logs.Counter, logText)
	fmt.Println("Задача отмечена как не выполненная - " + title)
}
