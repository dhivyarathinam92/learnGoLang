package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Model for course - file
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

// fake DB - created a slice
var courses []Course

// validation helper - file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers - file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to build API </h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses by unique Id")
	w.Header().Set("Content-Type", "application/json")
	// grab ID from req
	params := mux.Vars(r)
	// loop through slice and find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one courses")
	w.Header().Set("Content-Type", "application/json")
	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	// TODO check only if title is duplicate
	// loop , find out matching course name , JSON response

	// generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one courses")
	w.Header().Set("Content-Type", "application/json")
	// first grab id from req
	params := mux.Vars(r)
	// loop through slice and find matching id and remove the value and add with new ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one courses")
	w.Header().Set("Content-Type", "application/json")
	// first grab id from req
	params := mux.Vars(r)
	// loop through slice and find matching id and remove the value
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted Successfully")
			break
		}
	}
}

func main() {
	fmt.Println("API - Learn Go Code")
	r := mux.NewRouter()
	// injecting a data
	courses = append(courses, Course{CourseId: "2",
		CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{Fullname: "Author1", Website: "google1.com"}})

	courses = append(courses, Course{CourseId: "4",
		CourseName: "MernJS", CoursePrice: 199,
		Author: &Author{Fullname: "Author2", Website: "google2.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}
