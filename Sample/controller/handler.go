package controller

import (
	"Sample/greet"
	"log"
	"net/http"
)

func Handlers() {
	http.HandleFunc("/", greet.Helloworld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
