package Ucmds

import (
	"fmt"
	"to-do/logs"
)

func WrongInput(text []string) {
	logs.NewLog(logs.Counter, "Пользователь ввел неизвестную команду: "+text[0])
	fmt.Println("Вы ввели неизвестную команду")
	fmt.Println("Пожалуйста, введите команду, help покажет список команд")
}
