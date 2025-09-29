package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func (h*Handler)GetProductById(w http.ResponseWriter, r*http.Request){
	productId := r.PathValue("productId")
	fmt.Println("productid",productId)
	pId,err := strconv.Atoi(productId)
	fmt.Println("Type of productId:", reflect.TypeOf(productId))
	if err != nil{
		http.Error(w, "give valid ProductId", 400)
		return
	}

	product,err := h.productRepo.Get(pId)
	if err != nil{
		http.Error(w, "please give me a valid product id",400)
		return
	}
	if product == nil{
		util.SendError(w, 404, "produt not found")
	}
	util.SendData(w,product, 200)
}

