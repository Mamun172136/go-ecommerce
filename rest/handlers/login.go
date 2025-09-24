package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r*http.Request){
	var  reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil{
		http.Error(w,"invalid request data", 400)
	}

	usr := database.Find(reqLogin.Password,reqLogin.Email)
	if usr != nil{
		http.Error(w, "invalid credential",401)
	}

	util.SendData(w, usr, 201)
}