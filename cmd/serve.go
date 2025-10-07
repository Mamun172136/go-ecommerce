package cmd

import (
	"ecommerce/config"
	"ecommerce/infra/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"os"
)

func Serve() {

	cnf:= config.GetConfig()
	dbCon,err:= db.NewConnection(cnf.DB)
	if err!= nil{
		fmt.Println(err)
		os.Exit(0)		
	}
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	middlewares:= middleware.NewMiddlewares(cnf)

	productHandler :=  product.NewHandler(middlewares, productRepo)
	userHandler    :=  user.NewHandler(userRepo, cnf)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler, 

	)
	server.Start()
	// manager := middleware.NewManager()
	// manager.Use(middleware.Logger)
	// mux := http.NewServeMux()

	// // handler := http.HandlerFunc(handlers.Test)
	// initRoutes(mux, manager)
	// fmt.Println("Server running on port :8080") 

	// globalRouter := middleware.CorsWithPreflight(mux)
	// err := http.ListenAndServe(":8080", globalRouter)
	// if err != nil {
	// 	fmt.Println("error starting server", err)
	// }
}

// func init() {
// 	prd1 := database.Product{
// 		ID:          1,
// 		Title:       "Orange",
// 		Description: "Orange is red. I love orange.",
// 		Price:       100,
// 		ImgUrl:      "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
// 	}

// 	prd2 := database.Product{
// 		ID:          2,
// 		Title:       "Apple",
// 		Description: "Apple is green. I hate apple.",
// 		Price:       40,
// 		ImgUrl:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg",
// 	}

// 	prd3 := database.Product{
// 		ID:          3,
// 		Title:       "Banana",
// 		Description: "Banana is boring. I feel bored eating banana.",
// 		Price:       5,
// 		ImgUrl:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
// 	}

// 	// prd4 := Product{
// 	// 	ID: 4,
// 	// 	Title: "Angur Fol",
// 	// 	Description: "Angur Fol Tastes good.",
// 	// 	Price: 140,
// 	// 	ImgUrl: "https://cdn.dhakapost.com/media/imgAll/BG/2022January/angur-2-20220215152127.jpg",
// 	// }

// 	// prd5 := Product{
// 	// 	ID: 5,
// 	// 	Title: "Mango",
// 	// 	Description: "Mango is my favorite. I love it very much.",
// 	// 	Price: 1000000,
// 	// 	ImgUrl: "https://www.dole.com/sites/default/files/styles/512w384h-80/public/media/dole-blog-03-maerz-mango-05.jpg?itok=qXHJMEAz-PEthlz_-",
// 	// }

// 	// prd6 := Product{
// 	// 	ID: 6,
// 	// 	Title: "Strawberry",
// 	// 	Description: "Strawberries are sweet, juicy, and bursting with flavor.",
// 	// 	Price: 500,
// 	// 	ImgUrl: "https://snaped.fns.usda.gov/sites/default/files/styles/crop_ratio_7_5/public/seasonal-produce/2018-05/strawberries.jpg.webp?itok=B4LFd4vV",
// 	// }

// 	database.ProductList = append(database.ProductList, prd1)
// 	database.ProductList = append(database.ProductList, prd2)
// 	database.ProductList = append(database.ProductList, prd3)
// 	// productList = append(productList, prd4)
// 	// productList = append(productList, prd5)
// 	// productList = append(productList, prd6)
// }