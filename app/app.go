package main


import (
	"log"
	"net/http"
)


func start(){
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}