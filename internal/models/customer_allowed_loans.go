package models

type CustomerAllowedLoansResponse struct {
	Customer string `json:"customer"`
	Loans []Loan `json:"loans"`
}