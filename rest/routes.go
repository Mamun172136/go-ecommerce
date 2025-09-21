package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager){
	mux.Handle("/hello", manager.With(http.HandlerFunc(handlers.Test)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("POST /create-products", manager.With(http.HandlerFunc(handlers.CreateProducts)))

}