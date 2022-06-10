package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")//key value pairs

	languages := make(map[string]string)

	languages["JS"]="javascript"
	languages["RB"]="rubby"

	fmt.Println("list of languages: ",languages)
	fmt.Println("JS for ",languages["JS"])

	delete(languages,"RB")
	fmt.Println("list of languages: ",languages)

	//loops are interesting in go lang

	for key,value :=range languages{
		fmt.Printf("For key %v ,value is %v\n",key,value)
	}

	for _,value :=range languages{//key eka ona naththan 
		fmt.Printf("value is %v\n",value)
	}

}
