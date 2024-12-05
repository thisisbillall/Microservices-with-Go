package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Go!")
	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye Bye")
	})
	http.ListenAndServe(":9090", nil)

}
