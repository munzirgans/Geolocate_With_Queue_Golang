package main

import (
	"fmt"
	"golang_api_queue/sending/pkg/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/insertlocation", config.InsertLocation).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Terhubung dengan port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
