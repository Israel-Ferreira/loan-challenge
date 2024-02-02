package models

type Loan struct {
	Type LoanType `json:"type"`
	InterestRate float64 `json:"interest_rate"`
}


type LoanType string


const (
	Consignment LoanType = "CONSIGNMENT"
	Guaranteed = "GUARANTEED"
	Personal = "PERSONAL"
)


func NewPersonalLoan() Loan {
	return  Loan{
		Type: Personal,
		InterestRate: 4.0,
	}
}

func NewConsignmentLoan() Loan {
	return  Loan{
		Type: Consignment,
		InterestRate: 2.0,
	}
}

func NewGuaranteedLoan() Loan {
	return  Loan{
		Type: Guaranteed,
		InterestRate: 3.0,
	}
}