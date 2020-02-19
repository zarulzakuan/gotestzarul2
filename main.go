package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/thedevsaddam/gojsonq"
)

const (
	commentsURL = "https://jsonplaceholder.typicode.com/comments"
)

// Comments hold comments response
type Comments struct {
	PostID int    `json:"postId,omitempty"`
	ID     int    `json:"id,omitempty"`
	Email  string `json:"email,omitempty"`
	Body   string `json:"body,omitempty"`
	Title  string `json:"title,omitempty"`
}

func main() {
	// get comments count from all posts

	r := mux.NewRouter()
	r.HandleFunc("/comments/search/postid/{st}", searchPostID)
	r.HandleFunc("/comments/search/id/{st}", searchCommentID)
	r.HandleFunc("/comments/search/name/{st}", searchName)
	r.HandleFunc("/comments/search/email/{st}", searchEmail)
	r.HandleFunc("/comments/search/body/{st}", searchBody)

	log.Println("Running..")
	http.ListenAndServe(":8080", r)

}

func searchPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	st := vars["st"]
	sti, _ := strconv.Atoi(st)
	resp := getAllComments()
	var jsonn = fmt.Sprintf("{\"comments\":%s}", resp)
	avg := gojsonq.New().FromString(jsonn).From("comments")
	result := avg.Where("postId", "=", sti).Get()
	jsonString, _ := json.Marshal(result)
	w.Write(jsonString)
}
func searchCommentID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	st := vars["st"]
	sti, _ := strconv.Atoi(st)
	resp := getAllComments()
	var jsonn = fmt.Sprintf("{\"comments\":%s}", resp)
	avg := gojsonq.New().FromString(jsonn).From("comments")
	result := avg.Where("id", "=", sti).Get()
	jsonString, _ := json.Marshal(result)
	w.Write(jsonString)
}
func searchName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	st := vars["st"]
	resp := getAllComments()
	var jsonn = fmt.Sprintf("{\"comments\":%s}", resp)
	avg := gojsonq.New().FromString(jsonn).From("comments")
	result := avg.Where("name", "contains", st).Get()
	jsonString, _ := json.Marshal(result)
	w.Write(jsonString)
}
func searchEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	st := vars["st"]
	resp := getAllComments()
	var jsonn = fmt.Sprintf("{\"comments\":%s}", resp)
	avg := gojsonq.New().FromString(jsonn).From("comments")
	result := avg.Where("email", "contains", st).Get()
	jsonString, _ := json.Marshal(result)
	w.Write(jsonString)
}
func searchBody(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	st := vars["st"]
	resp := getAllComments()
	var jsonn = fmt.Sprintf("{\"comments\":%s}", resp)
	avg := gojsonq.New().FromString(jsonn).From("comments")
	result := avg.Where("body", "contains", st).Get()
	jsonString, _ := json.Marshal(result)
	w.Write(jsonString)
}

func getAllComments() string {
	// get all comments

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", commentsURL, nil)

	if err != nil {
		log.Println("Request failed")

	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("client do")
		log.Println(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}
