package cmd

import (
	"encoding/json"
	"github.com/Israel-Ferreira/loan-challenge/internal/handlers"
	"github.com/Israel-Ferreira/loan-challenge/internal/middlewares"
	"log"
	"net/http"
)

func StartServer(port string) {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		msg := map[string]string{
			"title":   "API Loans",
			"version": "0.1",
		}

		if err := json.NewEncoder(writer).Encode(msg); err != nil {
			errorMsg := map[string]string{
				"msg": err.Error(),
			}

			writer.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(writer).Encode(errorMsg)
			return
		}

	})

	http.HandleFunc(
		"/customer-loans",
		middlewares.Chain(handlers.CustomerLoansHandler, middlewares.LogMiddleware(), middlewares.JsonMiddleware()),
	)

	log.Fatalln(http.ListenAndServe(port, nil))
}
