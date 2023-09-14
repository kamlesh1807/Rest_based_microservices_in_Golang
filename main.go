package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Customer struct {
  
  Name string  `json:"full_name"`
  City string   `json:"city"`
  Zipcode string  `json:"zip_code"`
}

func main(){


	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomer)


	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}


func greet(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hello world ! \n how are you?")
}

func getAllCustomer(w http.ResponseWriter, r *http.Request){
 customers :=[]Customer {
	{Name:"aman", City:"vns", Zipcode:"221008"},
	{Name:"ayush", City:"delhi", Zipcode:"221778"},
	{Name:"piyush", City: "Noida", Zipcode: "110088"},

 }
   w.Header().Add("Content-Type", "application/json")
   json.NewEncoder(w).Encode(customers) //encoding the customers data as JSON and sending it as the response body to the client through the HTTP response.

}