package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler, 
	userHandler *user.Handler,

) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	
	}
}

func (server *Server) Start(){ 

	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	// handler := http.HandlerFunc(handlers.Test)
	// initRoutes(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	fmt.Println("Server running on port :3000")
	addr := ":" + strconv.Itoa(server.cnf.HttpPort)
	globalRouter := middleware.CorsWithPreflight(mux)
	err := http.ListenAndServe(addr, globalRouter)
	if err != nil {
		fmt.Println("error starting server", err)
	}

}