package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type reqCreateProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler)CreateProducts(w http.ResponseWriter, r *http.Request) {
	// handleCors(w)
	// handlePreflightReq(w,r)

	// if r.Method != "POST"{
	// 	http.Error(w, "plz give post request",400)
	// }

	var newProduct reqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "plz give valid json", 400)
		return
	}

	createdProduct, err:= h.productRepo.Create(repo.Product{
	Title: newProduct.Title,   
	Description :newProduct.Description,
	Price      :newProduct.Price,
	ImgUrl  : newProduct.ImgUrl ,   
	})
	if err != nil{
		http.Error(w,"Internal server error",http.StatusInternalServerError)
	}
	// createdProduct:= database.Store(newProduct)

	// newProduct.ID = len(database.ProductList) + 1

	// database.ProductList = append(database.ProductList, newProduct)

	// util.SendData(w, newProduct, 201)
	util.SendData(w, createdProduct, 201)
}
