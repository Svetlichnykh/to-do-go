package tasks

import (
	"fmt"
	"slices"
	"to-do/logs"
)

func HandleDel(text []string) {
	title := titleFrom(text)

	i, err := findTask(title)

	if err != nil {
		fmt.Println(err)
		logs.NewLog(logs.Counter, "Пользователь попылатался удалить задачу, но получил ошибку:")
		logs.NewLog(0, err.Error())
		return
	}

	Pool = slices.Delete(Pool, i, i+1)
	fmt.Println("Задача удалена - " + title)

}
