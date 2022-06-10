package main

import "fmt"

func main() {
	fmt.Println("Structs in go lang")

	//no inheritance in golang; no super or parent

	chamoda := User{"Chamoda","chamoda@gmail.com",true,22}
	fmt.Println(chamoda)
	fmt.Printf("Chamoda's details are: %+v\n",chamoda)//varibale name ekath ekka print wenne
	fmt.Printf("Chamoda's name is: %v\n",chamoda.Name)


}

type User struct {
	Name   string // first letter capital means - public
	Email  string
	status bool
	Age    int
}
