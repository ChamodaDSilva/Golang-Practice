package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //changing the json key name
	Price    int
	Platform string
	Password string   `json:"-"`              //hide value
	Tags     []string `json:"tags,omitempty"` //ignore nil fields
}

func main() {
	//EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 229, "Youtube", "abc123", []string{"web-dev", "js"}},
		{"Anguiler", 123, "Youtube", "adfa123", []string{"f-dev", "go"}},
		{"Mern", 22329, "Youtube", "2414dd", nil},
	}

	//package this data as json

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //data,prefix,suffix <-- parameters passing
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"Price": 229,
		"Platform": "Youtube",
		"tags": [
				"web-dev",
				"js"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)//show the data decoded
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	///// some cases where you just want to add data to key value pair//////

	var myOnlineData map[string]interface{}//key string and value is interface(when we dont know the returning data type value)
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)//get data from jsondatafrom web and add to myonlinedata
	fmt.Printf("%#v\n", myOnlineData)//printing the map

	for k, v := range myOnlineData {//showing separated data
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
