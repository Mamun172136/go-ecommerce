package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r*http.Request){
	productId := r.PathValue("productId")

	pId,err := strconv.Atoi(productId)
	if err != nil{
		http.Error(w, "give valid ProductId", 400)
		return
	}

	product := database.Get(pId)
	if product == nil{
		util.SendError(w, 404, "produt not found")
	}
	util.SendData(w,product, 200)
}

