package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf *config.Config){ 

	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	mux := http.NewServeMux()

	// handler := http.HandlerFunc(handlers.Test)
	initRoutes(mux, manager)
	fmt.Println("Server running on port :8080")
	addr := ":" + strconv.Itoa(cnf.HttpPort)
	globalRouter := middleware.CorsWithPreflight(mux)
	err := http.ListenAndServe(addr, globalRouter)
	if err != nil {
		fmt.Println("error starting server", err)
	}

}