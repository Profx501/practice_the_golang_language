/*
Dup выводит текст каждой строки, которая появляется в
стандартном вводе более одного раза, а так же кол-во
ее появлений.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("[INFO] ->> Кол-во: %d Слово: %s\n", n, line)
		}
	}
}
