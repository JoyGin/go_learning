package main

import (
	"fmt"
	"log"
	"net/http"
)
import "github.com/julienschmidt/httprouter"

// github.com/julienschmidt/httprouter

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_, err := fmt.Fprintf(w, "Welcome!\n")
	if err != nil {
		return
	}
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, err := fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	if err != nil {
		return
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
