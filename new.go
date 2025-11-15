package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your rollno: ")
	scanner.Scan()
	rollno := scanner.Text()

	fmt.Print("Enter your name: ")
	scanner.Scan()
	name := scanner.Text()

	var age int
	for {
		fmt.Print("Enter your age: ")
		scanner.Scan()
		text := scanner.Text()
		atoi, err := strconv.Atoi(text)
		age = atoi
		if err != nil {
			fmt.Println("Errors:", err)
		} else if atoi <= 0 {
			fmt.Println("Invilad ")
		} else {
			break
		}
	}
	var sex string
	for {
		fmt.Print("Enter your sex (male/female/other): ")
		scanner.Scan()
		text := scanner.Text()
		sex = text
		if text == "male" {
			break
		} else if text == "female" {
			break
		} else if text == "other" {
			break
		} else {
			fmt.Println(("Errors,please enter again"))
		}
	}

	var mark float64
	for {
		fmt.Print("Enter your desire PFC mark (0-10): ")
		scanner.Scan()
		text := scanner.Text()
		score, err := strconv.ParseFloat(text, 64)
		mark = score
		if err != nil {
			fmt.Println("Errors,please enter again your mark")
		} else if score < 0 || score > 10 {
			fmt.Println("Errors,please enter mark from 0 to 10")
		} else {
			break
		}
	}

	fmt.Println("\n============================")
	fmt.Println("Rollno:", rollno)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Sex:", sex)
	fmt.Println("Your desire PFC mark:", mark)
}
