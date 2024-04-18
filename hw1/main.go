// Требуется написать программу, которая по введенному полному пути выводит имя файла и расширение
// Напишите код, который выведет следующее
// filename: <name>
// extension: <extension>

// Подсказка. Возможно вам понадобится функция strings.LastIndex
// Для проверки своего решения используйте функции filepath.Base() filepath.Ext()
// Они могут помочь для проверки решения

// Сценарии
// ./filename-ext /home/robotomize/1.txt
// filename: 1
// extension: txt

// ./filename-ext /home/robotomize/1.txt.txt
// filename: 1.txt
// extension: txt

// ./filename-ext /home/robotomize/1
// filename: 1
// extension:

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := strings.Join(os.Args[1:], "") 					// filePth присваеваем значение из консоли, которое переведено в массив с удалением пробелов

	var fileName, fileExt string

	fullFileName := filePth[strings.LastIndex(filePth, `/`)+1:] // fullFileName присваеваем, в виде строки, значение последнего элемента массива filePth. (разделение по "/")
	sepFullName := strings.Split(fullFileName, ".")             // sepFullName присваеваем массив, который получили путем разбития fullFileName строки. (разбитие по ".")

	if len(sepFullName) >= 2 { 									// проверка на корректность названия файла с последующим присвоением итоговых значений
		fileName = sepFullName[0]
		fileExt = sepFullName[len(sepFullName)-1]
	} else if len(sepFullName) < 2 {
		fileName = sepFullName[0]
	}

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
