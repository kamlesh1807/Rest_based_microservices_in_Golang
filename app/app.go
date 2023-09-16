package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func Start(){
	//mux :=http.NewServeMux()
	router:= mux.NewRouter()
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomer)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer)
	router.HandleFunc("/create_customers", create_customer)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

