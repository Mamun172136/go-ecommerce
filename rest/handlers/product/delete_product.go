package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h*Handler)DeleteProduct(w http.ResponseWriter, r *http.Request){
	productId := r.PathValue("id")
	pId,err := strconv.Atoi(productId)
	if err != nil{
		http.Error(w, "please give me a valid product id",400)
		return
	}

	err = h.productRepo.Delete(pId)
	if err!= nil{
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	// database.Delete(pId)
	util.SendData(w,"successfully deleted product", 201)
}