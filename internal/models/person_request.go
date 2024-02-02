package models

type PersonRequest struct {
	Age uint `json:"age"`
	Cpf string `json:"cpf"`
	Name string `json:"name"`
	Income float64 `json:"income"`
	Location string `json:"location"`
}

