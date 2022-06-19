package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""

	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	//seeding - add data to slice
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Chamoda", Website: "Youtube"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Go", CoursePrice: 123, Author: &Author{Fullname: "Chamoda", Website: "Facebook"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Http", CoursePrice: 29129, Author: &Author{Fullname: "Chamoda", Website: "Youtube"}})

	//routing
	r.HandleFunc("/", serverHome).Methods("GET")// http://localhost:4000/
	r.HandleFunc("/courses", getAllCourses).Methods("GET")// http://localhost:4000/courses
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // http://localhost:4000/course/3
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", upddateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listern to a port
	log.Fatal(http.ListenAndServe(":4000", r)) // http://localhost:4000/

}

// controller -file

//server home route

func serverHome(w http.ResponseWriter, r *http.Request) { //r -reder get data , w writer send data/
	w.Write([]byte("<h1>Welcome to Api</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

//send the course id and check it is in or not
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//printing got parms
	for _, p := range params {
		fmt.Println(p)
	}

	//loop through courses and find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	//if not found
	var notFoundNotice = "No Course found in the given id" + params["id"]
	json.NewEncoder(w).Encode(notFoundNotice)
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what is :body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	//what about data not in the json-{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}

	//generate a unique id
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return

}

func upddateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//grab the request
	params := mux.Vars(r)

	//loop , id , remove , add my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //remove given id course from slice

			_ = json.NewDecoder(r.Body).Decode(&course)

			var course Course
			course.CourseId = params["id"]
			courses = append(courses, course)

		}
		//TODO: send a response when id is not found

	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	//grab the request
	params := mux.Vars(r)

	//loop , id , remove , add my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...) //remove given id course from slice

			json.NewEncoder(w).Encode("Deleted")
			break
		}
		//TODO: send a response when id is not found

	}
}
