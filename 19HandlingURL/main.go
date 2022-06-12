/* link wala thiyena parts check krna hati(port eka,schema eka wage ewa)*/
package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling urll in golang")
	fmt.Println(myurl)

	//parsing to url
	result, _ := url.Parse(myurl)

	/////////url eka danagena hityama parts walat wen krna widya
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of quesy params are : %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	///////////parts danagena hityama url eka hadanna puluwan widiyak
	parseOfUrl:= &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherURL := parseOfUrl.String()
	fmt.Println(anotherURL)
}
