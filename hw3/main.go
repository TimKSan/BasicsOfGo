package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type Item struct {
	link string
	name string
	tegs string
	date string
}

var items []Item

func assignment(args []string) Item {
	var it Item
	it.link = args[0]
	it.name = args[1]
	it.tegs = args[2]
	it.date = time.Now().Format(time.DateTime)
	items = append(items, it)
	return it
}

func main() {
	defer func() {
		// Завершаем работу с клавиатурой при выходе из функции
		_ = keyboard.Close()
	}()

	fmt.Println("Программа для добавления url в список")
	fmt.Println("Для выхода из приложения нажмите Esc")

OuterLoop:
	for {
		// Подключаем отслеживание нажатия клавиш
		if err := keyboard.Open(); err != nil {
			log.Fatal(err)
		}

		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		switch char {
		case 'a':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}

			// Добавление нового url в список хранения
			fmt.Println("Введите новую запись в формате <url описание теги>")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			args := strings.Fields(text)
			if len(args) < 3 {
				fmt.Println("Введите правильный аргументы в формате url описание теги")
				continue OuterLoop
			}
			// Напишите свой код здесь
			assignment(args)

		case 'l':
			// Вывод списка добавленных url. Выведите количество добавленных url и список с данными url
			// Вывод в формате
			// Имя: <Описание>
			// URL: <url>
			// Теги: <Теги>
			// Дата: <дата>

			// Напишите свой код здесь
			if len(items) != 0 {
				for i, item := range items {
					fmt.Printf("Кол-во добавленных url: %d\n %d. Имя: %s\n   URL: %s\n   Теги: %s\n   Дата: %s\n", len(items), i+1, item.name, item.link, item.tegs, item.date)
				}
			} else {
				fmt.Println("URL's is empty")
			}

		case 'r':
			if err := keyboard.Close(); err != nil {
				log.Fatal(err)
			}
			// Удаление url из списка хранения
			fmt.Println("Введите имя ссылки, которое нужно удалить")

			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			name := strings.TrimSpace(text)

			// Напишите свой код здесь
			for i, item := range items {
				if item.name == name {
					items = append(items[:i], items[i+1:]...)
					fmt.Println("ссылка удалена")
					break
				}
			}

		default:
			// Если нажата Esc выходим из приложения
			if key == keyboard.KeyEsc {
				return
			}
		}
	}
}
