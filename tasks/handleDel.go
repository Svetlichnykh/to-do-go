package tasks

import (
	"fmt"
	"slices"
	"strings"
	"to-do/logs"
)

func HandleDel(text []string) {
	var title string
	for i := range text {
		if i > 0 {
			title += text[i] + " "
		}
	}
	title = strings.TrimSpace(title)
	var logText string
	existFlag := false

	for i, v := range Pool {
		if strings.ToLower(v.title) == strings.ToLower(title) {
			existFlag = true
			Pool = slices.Delete(Pool, i, i+1)
			logText = "Пользователь удалил задачу - " + title
		}
	}
	if !existFlag {
		fmt.Println("Не найдено задачи с названием", title)
		logText = "Пользователь попытался удалить несуществующую задачу - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}
	logs.NewLog(logs.Counter, logText)
	fmt.Println("Задача удалена - " + title)
}
