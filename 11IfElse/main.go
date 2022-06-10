package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")

	var loginCount int = 23

	var result string

	if loginCount < 10 {
		result = "count down 10"
	}else if loginCount<20{
		result="count down 20 up 10"
	}else{
		result="something nmber"
	}

	fmt.Println(result)

	//////////////////////// initialize and checking at the same time

	if num := 3; num < 10 {
		result = "count down 10"
	}else{
		result="something nmber"
	}


}
