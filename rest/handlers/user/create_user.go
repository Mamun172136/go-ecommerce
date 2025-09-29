package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type reqCreateUser struct {
	ID          int    `json:"Id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler)CreateUser(w http.ResponseWriter, r*http.Request){

	var newUser reqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil{
		http.Error(w,"invalid data",400)
		return
	}

	createdUser,err := h.userRepo.Create(repo.User{
	FirstName: newUser.FirstName,
	LastName     : newUser.LastName,
	Email      :newUser.Email,
	Password     : newUser.Password,
	IsShopOwner  : newUser.IsShopOwner,
	})
	if err != nil{
		http.Error(w, "internal server error",http.StatusInternalServerError)
		return
	}
	
	util.SendData(w,createdUser,201)
}