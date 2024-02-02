package services

import "github.com/Israel-Ferreira/loan-challenge/internal/models"


const LOCATION_SAO_PAULO string =  "SP"

func ListCustomerLoansAllowed(customer models.PersonRequest) models.CustomerAllowedLoansResponse {
	
	customerName :=  customer.Name
	customerAge :=  customer.Age
	
	location := customer.Location
	
	
	var allowedLoans []models.Loan
	
	
	if (customer.Income <= 3000) || ((customer.Income > 3000 && customer.Income <=5000) && (customerAge < 30 && location == LOCATION_SAO_PAULO)) {
		allowedLoans =  append(allowedLoans, models.NewPersonalLoan())
	}
	
	
	
	if customer.Income >= 5000 {
		allowedLoans = append(allowedLoans, models.NewConsignmentLoan())
	}
	
	
	if customer.Income <= 3000 ||((customer.Income > 3000 && customer.Income <=5000) && (customerAge < 30 && location == LOCATION_SAO_PAULO)) {
		allowedLoans = append(allowedLoans, models.NewGuaranteedLoan())
	}
	
	
	
	return models.CustomerAllowedLoansResponse{
		Customer: customerName,
		Loans: allowedLoans,
	}
}