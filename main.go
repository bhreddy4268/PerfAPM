package main

import (
	"log"
	"net/http"

	"github.com/bhreddy4268/PerfAPM/apmservice"
)

func main() {

	apmservice.Init()
	log.Println("Starting server on port 4047")
	log.Println(http.ListenAndServe(":4047", nil))

}
