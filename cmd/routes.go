package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("/hello", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("/products", manager.With(http.HandlerFunc(handlers.GetProducts)))

	mux.Handle("/create-products", manager.With(http.HandlerFunc(handlers.CreateProducts)))

}