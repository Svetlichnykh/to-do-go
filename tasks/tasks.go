package tasks

import (
	"fmt"
	"time"
	"to-do/logs"
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
	PrintTask(0, newTask)

}
