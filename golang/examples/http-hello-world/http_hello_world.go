package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", mainPage)
	fmt.Println("Going to listen on http://localhost:5000  Ctr-c to stop the server.")
	err := http.ListenAndServe("127.0.0.1:5000", nil)
	// err := http.ListenAndServe("0.0.0.0:5000", nil)
	// err := http.ListenAndServe(":5000", nil) // defauls to 0.0.0.0 which is not secure
	if err != nil {
		log.Fatal(err)
	}
}
