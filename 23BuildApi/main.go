package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//model for course-file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake database -slice used
var courses []Course

//middleware -helper -file
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// controller -file

//server home route

func serverHome(w http.ResponseWriter, r *http.Request) { //r -reder get data , w writer send data/
	w.Write([]byte("<h1>Welcme to Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params:= mux.Vars(r)

	//printing got parms
	for _,p :=range params{
		fmt.Println(p)
	}

	//loop through courses and find matching id and return the response

	for _,course :=range courses{
		if course.CourseId ==params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	//if not found
	var notFoundNotice ="No Course found in the given id"+params["id"]
	json.NewEncoder(w).Encode(notFoundNotice)
}
