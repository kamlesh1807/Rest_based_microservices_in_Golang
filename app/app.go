package app

import (
	"log"
	"net/http"
	"github.com/kamlesh1807/Rest_based_microservices_in_Golang/service"
	"github.com/kamlesh1807/Rest_based_microservices_in_Golang/domain"

	"github.com/gorilla/mux"
)


func Start(){
	router:= mux.NewRouter()

	  ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}


	//mux :=http.NewServeMux()
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8080", router))

}

