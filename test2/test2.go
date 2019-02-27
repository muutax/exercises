package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello mazafaka")
}

func main() {
	http.HandleFunc("/", myHandler)
	panic(http.ListenAndServe(":8080", nil))
}
