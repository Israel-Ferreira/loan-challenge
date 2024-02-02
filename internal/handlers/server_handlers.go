package handlers

import (
	"encoding/json"
	"github.com/Israel-Ferreira/loan-challenge/internal/models"
	"github.com/Israel-Ferreira/loan-challenge/internal/services"
	"net/http"
)

func CustomerLoansHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}


	var personRequestBody models.PersonRequest

	if err := json.NewDecoder(request.Body).Decode(&personRequestBody); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	
	customerLoans :=  services.ListCustomerLoansAllowed(personRequestBody)
	
	if err := json.NewEncoder(writer).Encode(customerLoans); err !=  nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	
}