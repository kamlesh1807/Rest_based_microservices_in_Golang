package app

import (
	"encoding/json"
	"net/http"

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

   /* customers :=[]Customer {
	{Name:"aman", City:"vns", Zipcode:"221008"},
    }*/

    w.Header().Add("Content-Type", "application/json")
   //json.NewEncoder(w).Encode(customers) //encoding the customers data as JSON and sending it as the response body to the client through the HTTP response.
    json.NewEncoder(w).Encode(Customers)  ////encoding the customers data as XML and sending it as the response..

}