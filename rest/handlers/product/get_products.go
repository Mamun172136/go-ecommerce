package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler)GetProducts(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprint(w,"get product" )
	// if r.Method != "GET"{
	// 	http.Error(w,"plz give me get request",400)
	// }
	productList,err:= h.productRepo.List()
	if err != nil{
		http.Error(w, "please give me a valid product id",400)
		return
	}
	util.SendData(w,productList, 200)
}