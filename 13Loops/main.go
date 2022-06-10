package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop in go lang")

	days := []string{"sunday","monday","saturday","friday"}


	for d:=0;d<len(days);d++{
		fmt.Println(days[d])
	}

	for i:= range days{
		fmt.Println(days[i])
	}

	//////////////////////////
	for index ,day :=range days{
		fmt.Printf("index is %v and value is %v\n",index,day)
	}

	for _ ,day :=range days{//to ignore index
		fmt.Printf("value is %v\n",day)
	}

	//////////////////////////////////

	start :=1
	for start<10{
		if(start==5){
			start++
			continue
		}

		if(start==8){
			break
		}


		if(start==2){
			goto labelName//go to that label
		}

		fmt.Println(start)
		start++

	}
	/////////////////////////////////////

	labelName://label
		fmt.Println("Jumping to this label")
	


}
