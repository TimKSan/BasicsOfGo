// На основе шаблона напишите программу, подсчитывающую сколько раз буква встречается в предложении,
// а также частоту встречаемости в процентах. То есть в предложении hello world результатом будет:
// h - 1 0.1
// e - 1 0.1
// l - 3 0.33
// o - 2 0.2
// w - 1 0.1
// r - 1 0.1
// d - 1 0.1

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(strings.ToLower((strings.TrimSpace(text))), " ", "")

	letterCounts := make(map[string]int)

	for _, letter := range text {
		letterCounts[string(letter)]++
	}

	for k, v := range letterCounts {
		percent := float64(v) / float64(len(text)) * 100
		fmt.Printf("%s - %.2f\n", k, percent)
	}
}
