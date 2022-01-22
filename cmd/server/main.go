package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening at port 8020...")
	http.ListenAndServe(":8020", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{ \"message\": \"Hello world!\" }\n"))
}
