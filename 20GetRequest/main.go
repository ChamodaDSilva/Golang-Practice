package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video- LCO")

	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)

	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("COntent Length:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body) //content in the byte format

	var responseString strings.Builder //fmt.Println(string(content)) <-- meka wenuwata use krnne..dekenm ekama de wenne meke digai
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count is", byteCount)
	fmt.Println("Response body is", responseString.String())

	// fmt.Println(string(content))

}
