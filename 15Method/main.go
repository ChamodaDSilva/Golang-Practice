//functions of structs are methods
package main

import "fmt"

func main() {
	fmt.Println("methods")
	chamoda := User{"Chamoda","chamoda@gmail.com",true,22}
	
	

	chamoda.GetStatus()

	chamoda.email="newemail"
	fmt.Println(chamoda)

}

type User struct {
	Name   string // first letter capital means - public
	email  string
	Status bool
	Age    int
	
}

func(u User) GetStatus(){
	fmt.Println("Is the user active: ",u.Status)
}

func(u User) setEmail(){
	u.email="newemail.gmail"
}

//functions of structs are methods