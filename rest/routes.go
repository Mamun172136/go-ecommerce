package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("/hello", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("/products", manager.With(http.HandlerFunc(handlers.GetProducts)))

	mux.Handle("/create-products", manager.With(http.HandlerFunc(handlers.CreateProducts)))

}