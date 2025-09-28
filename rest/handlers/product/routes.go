package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler)RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager){

	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductById)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProducts)))


}