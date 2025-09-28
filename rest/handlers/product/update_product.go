package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h*Handler)UpdateProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("productId")
	pId,err := strconv.Atoi(productId)
	if err != nil{
		http.Error(w, "please give me a valid product id",400)
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil{
		http.Error(w, "please give me valid json",400)
		return
	}
	
	newProduct.ID = pId
	database.Update(newProduct)
	util.SendData(w,"successfully updated product", 201)
}