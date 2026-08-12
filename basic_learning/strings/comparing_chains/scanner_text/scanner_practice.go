package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Add some text: ")
	if !scanner.Scan() {
		return // stdin closed
	}
	command := strings.TrimSpace(scanner.Text())

	fmt.Println("> You added: ", command)
}
