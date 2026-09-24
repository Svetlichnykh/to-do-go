package tasks

import (
	"fmt"
	"strings"
	"time"
	"to-do/logs"
)

func HandleDone(text []string) {
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
			if v.done == true {
				fmt.Println("Данная задача уже отмечена как выполненная - " + title)
				logText = "Пользователь попытался отметить выполненную задачу выполненной - " + title
				logs.NewLog(logs.Counter, logText)
				return
			}
			v.doneDate = time.Now()
			v.done = true
			logText = "Пользователь отметил следующую задачу как выполненную - " + title
		}
	}
	if !existFlag {
		fmt.Println("Не найдено задачи с названием", title)
		logText = "Пользователь попытался отметить несуществующую задачу как выполненную - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}
	logs.NewLog(logs.Counter, logText)
	fmt.Println("Задача выполнена - " + title)
}
