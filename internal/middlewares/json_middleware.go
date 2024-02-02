package middlewares

import (
	"log"
	"net/http"
)

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

func LogMiddleware() Middleware {
	return func(handlerFunc http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			
			defer func() {
				log.Default().Println(request.URL)
				log.Println("Metodo Utilizado: ", request.Method)
				log.Println("Content-Type da Requisição: ", request.Header.Get("Content-Type"))
			}()
	
			handlerFunc(writer, request)
		}
	}
}


func JsonMiddleware() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Content-Type", "application/json")
			f(writer, request)
		}
	}
}


func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f =  m(f)
	}
	
	return f
}