package app

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

type Customer struct {
  
	Name string  `json:"full_name" xml:"name"` 
	City string   `json:"city" xml:"city"`
	Zipcode string  `json:"zip_code" xml:"zipcode"`
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

    w.Header().Add("Content-Type", "application/xml")
   //json.NewEncoder(w).Encode(customers) //encoding the customers data as JSON and sending it as the response body to the client through the HTTP response.
    xml.NewEncoder(w).Encode(customers)  ////encoding the customers data as XML and sending it as the response..

}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars:=mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func create_customer(w http.ResponseWriter, r *http.Request) {

fmt.Fprint(w, "Post received ")

}