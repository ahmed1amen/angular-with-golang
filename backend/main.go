package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)
type Course struct{
    Title string `json:"Title"`
    Dec string `json:"desc"`
    Content string `json:"content"`
}
type Courses []Course

func allCourses(w http.ResponseWriter ,r *http.Request){
w.Header().Set("Content-Type", "application/json")
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Headers", "Cache-Control, Pragma, Origin, Authorization, Content-Type, X-Requested-With")
w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST")
articles := Courses{
 Course{Title: "Hello", Dec: "Course Description", Content: "Course Content"},
 Course{Title: "Hello 2", Dec: "Course Description", Content: "Course Content"},
}

fmt.Println("Endpoint: Hit All Course")
json.NewEncoder(w).Encode(articles)

}

func homePage(w http.ResponseWriter ,r *http.Request)  {
	fmt.Fprint(w,"HomePage Endpoint HIT")
}

func handleRequest()  {

	http.HandleFunc("/",homePage)
	http.HandleFunc("/courses",allCourses)
	log.Fatal(http.ListenAndServe(":8000",nil))
}

func main()  {
	handleRequest()
	
}