package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome files in go lang")

	content := "this string should go to the file......"

	file, err:= os.Create("./myLocalFile.txt")

	// if err != nil{
	// 	panic(err)//shutdown the execution
	// }
	checkNillErr(err)


	length,err:=io.WriteString(file,content)

	// if err != nil{
	// 	panic(err)//shutdown the execution
	// }

	checkNillErr(err)

	fmt.Println("Length is",length)

	defer file.Close()


	readfile("./myLocalFile.txt")
}

func readfile(filename string){
	dataByte,err:=ioutil.ReadFile(filename)

	// if err != nil{
	// 	panic(err)//shutdown the execution
	// }
	checkNillErr(err)

	fmt.Println("text data inside the file is \n",string(dataByte))

}

func checkNillErr(err error){
	if err != nil{
		panic(err)//shutdown the execution
	}

}
