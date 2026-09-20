package main

import (
	"fmt"
	"to-do/input"

	"github.com/k0kubun/pp"
)

func main() {

	fmt.Println("========== HELLO USER ==========")
	fmt.Println("==== WELCOME TO MY TODO APP ====")
	hollowsCount := 0

	for {

		fmt.Print("> ")
		if ok := input.Scanner.Scan(); !ok {
			fmt.Println("Ошибка")
			return
		}

		text := input.Scanner.Text()
		if text == "" {

			if hollowsCount >= 3 {
				pp.Println("Идиот, хватит тыкать ENTER")
			}

			fmt.Println("Пожалуйста, введите команду")
			hollowsCount++

		} else {

			hollowsCount = 0
			fmt.Println("Ваша команда:", text)

			if text == "add" {
				fmt.Println("Добавляю задачу...")
			}

			if text == "exit" {
				fmt.Println("ВЫХОД...")
				return
			}

		}

	}
}
