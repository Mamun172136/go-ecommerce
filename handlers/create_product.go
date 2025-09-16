package handlers

import (
	
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w,r)

	// if r.Method != "POST"{
	// 	http.Error(w, "plz give post request",400)
	// }

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give valid json", 400)
		return
	}

	newProduct.ID = len(database.ProductList) + 1

	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
}
