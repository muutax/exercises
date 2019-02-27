package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("top_kek.txt")
	if err != nil {
		log.Fatal("unable to read from file: ", err)
	}
	fmt.Fprint(w, string(file))
}

func main() {
	var url string
	fmt.Scan(&url)
	resp, err := http.Get("https://" + url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
	http.HandleFunc("/", myHandler)
	file, err := os.Create("top_kek.txt")
	if err != nil { // если возникла ошибка
		log.Fatal("unable to create file: ", err)
	}
	/*

	 */

	fmt.Fprint(file, string(body))
	file.Close()

	panic(http.ListenAndServe(":8080", nil))
}
