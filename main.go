package main

import (
	"fmt"
)

func main() {
	op := ""
	left := 0
	right := 0
	guide := ""
	isOver := false
	yORn := ""

	fmt.Println("Welcome to the CLI calculator. Type 'help' for guide or 'go' to continue")
	fmt.Scan(&guide)

	if guide == "help" {
		fmt.Println(".......................")
		fmt.Println(" ")
		fmt.Println("Available operations are: 'add', 'mul', 'sub', 'div'")
		fmt.Println("Example Usage: add <num> <num> then click enter")
		fmt.Println(".......................")
		fmt.Println(" ")
		main()
	} else {
		fmt.Println("Type your operation")

		for isOver != true {
			fmt.Scan(&op, &left, &right)
			// if left >= 0 || right <= 9 {
			// 	fmt.Println("Usage Example: add <num> <num>. Give an operation followed by two numbers (add 1 2)")
			// 	continue
			// }

			switch op {
			case "add" :
				fmt.Println(add(left, right));
			case "sub" :
				fmt.Println(sub(left, right));
			case "mul" :
				fmt.Println(mul(left, right));
			case "div" :
				if right == 0 {
		        fmt.Println("Cannot divide by zero")
				continue
	        }
				fmt.Println(div(left, right));
			}

			fmt.Println("Do you wish to continue [yes/quit]")
			fmt.Println(".......................")
		    fmt.Println(" ")
			fmt.Scan(&yORn)

			if yORn == "quit" {
				isOver = true
				continue
			} else {
				fmt.Println("Input your operation")
			}
		}
	}
}

func add(left, right int) int {
	return left + right
}

func sub(left, right int) int {
	return left - right
}

func mul(left, right int) int {
	return left * right
}

func div(left, right int) int {
	return left / right
}