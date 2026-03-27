package main

import (
	"fmt"
	"strconv"
	"strings"
	// "strings"
)

func main() {
	arg := ""
	num := ""
	op := ""
	isOver := false
	yORn := ""

	fmt.Println("Welcome to the base converter. Example Usage: convert 1E hex")
	for isOver != true {
		fmt.Scan(&arg, &num, &op)
		fmt.Println(" ")

		if op == "hex" {fmt.Print("Decimal:"); fmt.Print(" ") ;fmt.Println(hex(num)); fmt.Println(" ")}
		if op == "bin" {fmt.Print("Decimal:"); fmt.Print(" ") ;fmt.Println(bin(num)); fmt.Println(" ")}
		if op == "dec" {dec(num); fmt.Println(" ")}

		fmt.Println("Do you wish to continue [yes/quit]")
		fmt.Println(".......................")
		fmt.Scan(&yORn)

		if yORn == "quit" {
			isOver = true
			continue
		} else {
			fmt.Println("Input your conversion")
		}
	}
}

func hex(hexStr string) (int64) {
	toDec, _ := strconv.ParseInt(hexStr, 16, 64)
	return toDec
}

func bin(binStr string) (int64) {
	toDec, _ := strconv.ParseInt(binStr, 2, 64)
	return toDec
}

func dec(decStr string) {
	decInt, _ := strconv.ParseInt(decStr, 10, 64)
	toBin := strconv.FormatInt(decInt, 2)
	toHex := strconv.FormatInt(decInt, 16)

	fmt.Println("Binary: ", toBin)
	fmt.Println("Hex: ", strings.ToUpper(toHex))
}