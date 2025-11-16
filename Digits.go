package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var number int
	for {
		fmt.Print("Enter your number: ")
		scanner.Scan()
		text := scanner.Text()
		input, err := strconv.Atoi(text)
		number = input
		if err != nil {
			fmt.Println("Errors,please enter number: ")
		} else {
			break
		}
	}

	if number < 0 {
		number = -number
	} else if number == 0 {
		fmt.Println("Digits: 0 ")
		return
	}
	fmt.Print("Digits: ")
	for number > 0 {
		digits := number % 10
		number = number / 10
		fmt.Print(digits)
		if number > 0 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
