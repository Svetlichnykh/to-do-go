package Ucmds

import (
	"fmt"
	"to-do/logs"
)

func Help(text []string) {
	if len(text) > 1 {
		logs.NewLog(logs.Counter, "Пользователь ввел help с лишними аргументами")
		fmt.Println("Команда help не принимает аргументов, введите help без аргументов для вывода доступных команд")
	} else {
		logs.NewLog(logs.Counter, "Пользователь вывел список команд")
		fmt.Println("Ваша команда - help")
		fmt.Println("")
		fmt.Println("Список команд:")
		fmt.Println("")
		fmt.Println("help — эта команда позволяет узнать доступные команды и их формат")
		fmt.Println("")
		fmt.Println("add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач")
		fmt.Println("")
		fmt.Println("list — эта команда позволяет получить полный список всех задач")
		fmt.Println("")
		fmt.Println("del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку")
		fmt.Println("")
		fmt.Println("done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную")
		fmt.Println("")
		fmt.Println("log — эта команда позволяет получить лог текущего выполнения")
		fmt.Println("")
		fmt.Println("exit — эта команда позволяет завершить выполнение программы")
		fmt.Println("")
	}
}
