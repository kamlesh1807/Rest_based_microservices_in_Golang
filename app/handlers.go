package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kamlesh1807/Rest_based_microservices_in_Golang/service"
)

type Customer struct {
  
	Name string  `json:"full_name" xml:"name"` 
	City string   `json:"city" xml:"city"`
	Zipcode string  `json:"zip_code" xml:"zipcode"`
  }

  type CustomerHandlers struct{
	service service.CutomerService
  }

  func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request){

	Customers, _:=ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")     //encoding the customers data as JSON.
    json.NewEncoder(w).Encode(Customers)  

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	id := vars["customer_id"]

	customer,err:=ch.service.GetCustomer(id)
	if err!=nil{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	}else{
		w.Header().Add("Content-Type","application/json")
		json.NewEncoder(w).Encode(customer)
	}
	
}
