package product

import (
	"ecommerce/repo"
	"ecommerce/rest/middleware"
)

type Handler struct {
	productRepo repo.ProductRepo
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middlewares, productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}