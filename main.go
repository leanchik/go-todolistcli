package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var tasks []string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню")
		fmt.Println("1. Добавить задачу")
		fmt.Println("2. Показать задачи")
		fmt.Println("3. Удалить задачу")
		fmt.Println("4. Очистить весь список")
		fmt.Println("5. Выход")
		fmt.Print("Выберите действие")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("Введите задачу")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			fmt.Println("Задача добавлена!")

		case "2":
			if len(tasks) == 0 {
				fmt.Println("Список задач пуст")
			} else {
				fmt.Println("Ваши задачи:")
				for i, task := range tasks {
					fmt.Printf("%d. %s\n", i+1, task)
				}
			}

		case "3":
			if len(tasks) == 0 {
				fmt.Println("Удалять нечего, список пуст")
				continue
			}
			fmt.Println("Введите номер задачи для удаления")
			numStr, _ := reader.ReadString('\n')
			numStr = strings.TrimSpace(numStr)
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Некорректный номер задачи")
				continue
			}
			tasks = append(tasks[:num-1], tasks[num:]...)
			fmt.Println("Задача удалена")

		case "4":
			tasks = []string{}
			fmt.Println("Список задач очищен!")

		case "5":
			fmt.Println("Пока!")
			return

		default:
			fmt.Println("Неизвестная команда. Введите 1-5.")
		}
	}
}
