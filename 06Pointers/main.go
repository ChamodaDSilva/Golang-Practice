package main

import "fmt"

func main() {
	fmt.Println("WELCMOME TO A CLASS ON POINTERS")

	// var pointer *int
	// fmt.Println(pointer)

	myNumber :=23

	var pointer2 =&myNumber

	fmt.Println("Value of",pointer2)
	fmt.Println("actual value", *pointer2)

	*pointer2 =*pointer2 + 3
	fmt.Println("Value of my number now",myNumber)

}
