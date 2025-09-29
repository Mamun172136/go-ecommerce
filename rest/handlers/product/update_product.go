package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqCreateProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h*Handler)UpdateProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("productId")
	pId,err := strconv.Atoi(productId)
	if err != nil{
		http.Error(w, "please give me a valid product id",400)
		return
	}

	var newProduct ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil{
		http.Error(w, "please give me valid json",400)
		return
	}
	
	newProduct.ID = pId 
	_, err= h.productRepo.Update(repo.Product{
	ID: pId,
	Title: newProduct.Title,   
	Description :newProduct.Description,
	Price      :newProduct.Price,
	ImgUrl  : newProduct.ImgUrl ,   
	})

	if err != nil{
		http.Error(w,"internal server error", http.StatusInternalServerError)
	}
	util.SendData(w,"successfully updated product", 201)
}