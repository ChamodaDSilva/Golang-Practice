package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to functions in golang")

	greeter()

	result := adder(3, 5)
	fmt.Println("result is", result)

	resultMore := proAdder(3, 5, 3, 4, 5)
	fmt.Println("result is", resultMore)

	number, text := returnMultipleValues(1, 2)
	fmt.Println(number, text)

}

func proAdder(values ...int) int { //get unlimited parametes and making slice named values
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func returnMultipleValues(num1 int, num2 int) (int, string) {//return multiple values
	return num1 + num2, "hiii"
}

func adder(num1 int, num2 int) int {
	return num1 + num2
}



func greeter() {
	fmt.Println("Hiii users")
}
