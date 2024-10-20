package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input sentences: ")
	words, _ := reader.ReadString('\n')

	words = strings.TrimSpace(words)

	fmt.Println(reverseString(words))
}

func reverseString(s string) string {
	arrStr := strings.Split(s, "")
	var newArr []string

	for i := len(arrStr) - 1; i >= 0; i-- {
		newArr = append(newArr, arrStr[i])
	}

	return strings.Join(newArr, "")
}
