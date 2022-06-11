package main

import "fmt"

func main() {
	defer fmt.Println("hello world")// defer dammama e line eka function eke anthimata run wenne
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Welcome to defer")
	
	mydefer()
}

func mydefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

//used to delay the execution of a function or a statement until the nearby function returns. In simple words, defer will move the execution of the statement to the very end inside a function.