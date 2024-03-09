/*
Echo выводит аргументы командной строки
Запуск: go run echo1.go Hello world!
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Вывод имени выполняемой программы
	fmt.Println(os.Args[0])
	// Вывод аргументов командной строки
	fmt.Println(strings.Join(os.Args[1:], " "))
}
