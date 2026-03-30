// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Hamza Musa]
// Squad:  [The Interfaces]

package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
// Your program accepts a command followed by a string and applies the correct transformation.

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("Type your command and a string")
	line, _ :=  reader.ReadString('\n')
	line = strings.TrimSpace(line)

	parts := strings.Fields(line)

	command := parts [0]
	textSlice := parts[1:]
	text := strings.Join(textSlice, " ")

	if command == "upper" {fmt.Println(upper(text))}
}

func upper(s string) string {
	return strings.ToUpper(s)
}
// func lower(s string) string {}
// func cap(s string) string {}
// func title(s string) string {}
// func snake(s string) string {}
// func reverse(s string) string {}