package tasks

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"to-do/logs"
	"to-do/timecalc"
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
			fmt.Print(" ( " + timecalc.TimeUntil(v.targetDate) + " )")
		}
		fmt.Println("")
	}
	if !v.doneDate.IsZero() {
		fmt.Println("✅ Время выполнения:", v.doneDate.Format("2006.01.02 15:04"))
	}
}

func EditTask(newTask Task, title string) {
	existFlag := false
	var logText string
	var changesExist bool
	var changes string
	var changedTaskId int
	for i := range Pool {
		v := &Pool[i]
		if strings.ToLower(v.title) == strings.ToLower(title) {
			existFlag = true
			if newTask.title != "" {
				changes += "Название: " + v.title + " --> " + newTask.title + "\n"
				logText += "Название: " + v.title + " --> " + newTask.title + "\n"
				v.title = newTask.title
				changesExist = true
			}
			if newTask.description != "" {
				if v.description != "" {
					changes += "Описание: " + v.description + " --> " + newTask.description + "\n"
					logText += "Описание: " + v.description + " --> " + newTask.description + "\n"
				} else {
					changes += "Описание: " + "+++ " + newTask.description + "\n"
					logText += "Описание: " + "+++ " + newTask.description + "\n"
				}
				v.description = newTask.description
				changesExist = true
			}
			if newTask.category != "" {
				if v.category != "" {
					changes += "Категория: " + v.category + " --> " + newTask.category + "\n"
					logText += "Категория: " + v.category + " --> " + newTask.category + "\n"
				} else {
					changes += "Категория: " + "+++ " + newTask.category + "\n"
					logText += "Категория: " + "+++ " + newTask.category + "\n"
				}
				v.category = newTask.category
				changesExist = true
			}
			if !newTask.creationDate.IsZero() {
				if !v.targetDate.IsZero() && newTask.creationDate.After(v.targetDate) {
					fmt.Println("Время создания не валидно, оно не может быть после дедлайна. Дата создания изменена не будет")
				} else if !v.doneDate.IsZero() && newTask.creationDate.After(v.doneDate) {
					fmt.Println("Время создания не валидно, оно не может быть после даты выполнения. Дата создания изменена не будет")
				} else {
					changes += "Дата создания: " + v.creationDate.Format("2006.01.02 15:04") + " --> " + newTask.creationDate.Format("2006.01.02 15:04") + "\n"
					logText += "Дата создания: " + v.creationDate.Format("2006.01.02 15:04") + " --> " + newTask.creationDate.Format("2006.01.02 15:04") + "\n"
					v.creationDate = newTask.creationDate
					changesExist = true
				}

			}
			if !newTask.targetDate.IsZero() {
				if !v.creationDate.IsZero() && newTask.targetDate.Before(v.creationDate) {
					fmt.Println("Время дедлайна не валидно, оно не может быть до даты создания. Дата дедлайна изменена не будет")
				} else {
					if !v.targetDate.IsZero() {
						changes += "Дедлайн: " + v.targetDate.Format("2006.01.02 15:04") + " --> " + newTask.targetDate.Format("2006.01.02 15:04") + "\n"
						logText += "Дедлайн: " + v.targetDate.Format("2006.01.02 15:04") + " --> " + newTask.targetDate.Format("2006.01.02 15:04") + "\n"
					} else {
						changes += "Дедлайн: " + "+++ " + newTask.targetDate.Format("2006.01.02 15:04") + "\n"
						logText += "Дедлайн: " + "+++ " + newTask.targetDate.Format("2006.01.02 15:04") + "\n"
					}
					v.targetDate = newTask.targetDate
					changesExist = true
				}
			}
			if !newTask.doneDate.IsZero() {
				if !v.creationDate.IsZero() && newTask.doneDate.Before(v.creationDate) {
					fmt.Println("Время выполнения не валидно, оно не может быть до даты создания. Дата выполнения изменена не будет")
				} else {
					if !v.doneDate.IsZero() {
						changes += "Дата выполнения: " + v.doneDate.Format("2006.01.02 15:04") + " --> " + newTask.doneDate.Format("2006.01.02 15:04") + "\n"
						logText += "Дата выполнения: " + v.doneDate.Format("2006.01.02 15:04") + " --> " + newTask.doneDate.Format("2006.01.02 15:04") + "\n"
					} else {
						changes += "Дата выполнения (теперь задача отмечена как выполненная): " + "+++ " + newTask.doneDate.Format("2006.01.02 15:04") + "\n"
						logText += "Дата выполнения (теперь задача отмечена как выполненная): " + "+++ " + newTask.doneDate.Format("2006.01.02 15:04") + "\n"
						v.done = true
					}
					v.doneDate = newTask.doneDate
					changesExist = true
				}
			}
			changedTaskId = i
		}
	}

	if !existFlag {
		fmt.Println("Не найдено задачи с названием", title)
		logText = "Пользователь попытался изменить несуществующую задачу - " + title
		logs.NewLog(logs.Counter, logText)
		return
	}

	if !changesExist {
		fmt.Println("Ничего не изменено, причины написаны выше")
		logText += "Пользователь совершил ошибки при заполнении дат и ничего не изменил в задаче - " + title
	} else {
		logText = "Пользователь изменил задачу - " + title + "\n" + "Изменения:" + "\n" + logText + "\n"
		fmt.Println("Вы успешно изменили задачу - " + title)
		fmt.Println("Изменения:")
		fmt.Println(changes)
		fmt.Println("Измененная задача:")
		PrintTask(-1, Pool[changedTaskId])
	}

	logs.NewLog(logs.Counter, logText)
}
