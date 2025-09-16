package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w,"get product" )
	// if r.Method != "GET"{
	// 	http.Error(w,"plz give me get request",400)
	// }

	util.SendData(w, database.ProductList, 200)
}