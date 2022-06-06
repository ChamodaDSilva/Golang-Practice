package main

import (
	"fmt"

)

const LoginToken string ="jahdakka" //if first letter is capital it tells that is capital
var outsider int =8

func main() {
	var username string = "chamoda"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var smallVal uint8 = 12 //only very small values
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var val int = 1231
	fmt.Println(val)
	fmt.Printf("Variable is of type: %T \n", val)

	var smallFloat float32 = 1.21 //only very small values
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var floVal float64 = 112131.21
	fmt.Println(floVal)
	fmt.Printf("Variable is of type: %T \n", floVal)

	//default values-----------------

	var anotherVariable int
	fmt.Println("Val :",anotherVariable)

	//implicit type---------------------------
	var web="google.io"
	fmt.Println(web)//the system decide the type

	//no var style----------------------------------
	numberOfUsers := 3000
	fmt.Println(numberOfUsers)

	fmt.Println(outsider)


}
