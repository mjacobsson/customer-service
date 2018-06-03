package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type customer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetCustomers GET /customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

// GetCustomer GET /customers/{id}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestID := params["id"]
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer{ID: requestID, Name: "Customer_" + requestID})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", GetCustomer).Methods("GET")
	if err := http.ListenAndServe(":8005", router); err != nil {
		log.Fatal(err)
	}
}
