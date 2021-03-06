package main

import (
	"log"
	"net/http"
	"sync"
)

var (
	mtx sync.RWMutex
)

func main() {

	http.HandleFunc("/join", Join)
	http.HandleFunc("/leave", Leave)
	http.HandleFunc("/refill", Refill)
	http.HandleFunc("/list", List)
	http.HandleFunc("/reset", Reset)
	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/config", ConfigUpdate)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
