package main

import (
	"github.com/Israel-Ferreira/loan-challenge/cmd"
	"log"
)


func main() {
	port := ":8000"
	log.Default().Println("Iniciando o Servidor na Porta 8000")
	cmd.StartServer(port)
}

