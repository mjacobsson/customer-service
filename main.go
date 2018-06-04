package main

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var customers = []customer{
	customer{ID: 1, Name: "Customer_1"},
	customer{ID: 2, Name: "Customer_2"},
	customer{ID: 3, Name: "Customer_3"},
}

// GetCustomers GET /customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// GetCustomer GET /customers/{id}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	requestID, _ := strconv.Atoi(params["id"])
	for _, customer := range customers {
		if customer.ID == requestID {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(customer)

			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", GetCustomer).Methods("GET")
	if err := http.ListenAndServe(":8005", router); err != nil {
		log.Fatal(err)
	}
}
