/* this code gives the htm code of the given website */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.cricbuzz.com/" //const kiyanne java final kiyana eka, change krnn be

func main() {
	fmt.Println("web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Body of request", string(content))

}
