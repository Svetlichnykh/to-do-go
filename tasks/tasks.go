package tasks

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
	"to-do/logs"
	"to-do/timeCalc"
)

type Task struct {
	title        string
	description  string
	category     string
	done         bool
	creationDate time.Time
	targetDate   time.Time
	doneDate     time.Time
}

type Options struct {
	Description string
	Category    string
	TargetDate  time.Time
}

var Pool = make([]Task, 0)

func NewTask(title string, options Options) {
	newTask := Task{
		title:        title,
		description:  options.Description,
		category:     options.Category,
		done:         false,
		creationDate: time.Now(),
		targetDate:   options.TargetDate,
	}
	Pool = append(Pool, newTask)
	logs.NewLog(logs.Counter, "Пользователь добавил задачу:")

	logTask := "Название: " + title + "\n"
	if newTask.description != "" {
		logTask += "Описание: " + newTask.description + "\n"
	}
	if newTask.category != "" {
		logTask += "Категория: " + newTask.category + "\n"
	}

	logTask += "Время создания: " + newTask.creationDate.Format("2006.01.02 15:04") + "\n"

	if !newTask.targetDate.IsZero() {
		logTask += "Дедлайн: " + newTask.targetDate.Format("2006.01.02 15:04") + "\n"
	}
	logs.NewLog(0, logTask)

	fmt.Println("Задача добавлена:")
	PrintTask(-1, newTask)

}

func PrintTask(i int, v Task) {
	if i != -1 {
		fmt.Print(strconv.Itoa(i+1) + ". ")
	}

	if v.category != "" {
		fmt.Print("(", v.category, ") ")
	}
	fmt.Print(v.title)
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
		fmt.Print("❗ Дедлайн: ", v.targetDate.Format("2006.01.02 15:04"))
		if !v.done {
			fmt.Print(" ( " + timeCalc.TimeUntil(v.targetDate) + " )")
		}
		fmt.Println("")
	}
	if !v.doneDate.IsZero() {
		fmt.Println("✅ Время выполнения:", v.doneDate.Format("2006.01.02 15:04"))
	}
	fmt.Println("")
}

func EditTask(newTask Task, title string) {
	for i, _ := range Pool {
		v := &Pool[i]
		existFlag := false
		var logText string
		if strings.ToLower(v.title) == strings.ToLower(title) {
			existFlag = true
			Pool = slices.Delete(Pool, i, i+1)
			logText = "Пользователь удалил задачу - " + title
		}
		if existFlag {
		}
		logs.NewLog(logs.Counter, logText)
	}
}
