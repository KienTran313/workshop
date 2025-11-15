package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var number int
	var err error
	for {
		fmt.Print("Enter your number here: ")
		fmt.Scan(&input)

		// Chuyển đổi chuỗi thành số nguyên
		number, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error,please enter again")
			continue
		}
		break
	}

	if number < 0 {
		number = -number
	}

	if number == 0 {
		fmt.Println("Digits: 0")
		return
	}

	var digits []int
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}

	fmt.Print("Digits: ")
	for i := len(digits) - 1; i >= 0; i-- {
		fmt.Print(digits[i])
		if i != 0 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
