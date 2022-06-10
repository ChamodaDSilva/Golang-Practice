package main

import "fmt"

func main() {
	fmt.Println("welcome to array in go lang")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[2] = "orange"
	fruitlist[3] = "peach"
	
	fmt.Println("Fruit list: ",fruitlist)
	fmt.Println("Fruit list size: ",len(fruitlist))

	var vegList=[5]string {"potato","beans","mushroom"}
	fmt.Println("Fruit list size: ",len(vegList))

}
