package tasks

import (
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
	logs.NewLog(logs.Counter, "Пользователь добавил задачу с названием "+newTask.title)
}
