package user

import (
	"ecommerce/config"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h*Handler) Login(w http.ResponseWriter, r*http.Request){
	var  reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil{
		http.Error(w,"invalid request data", 400)
		return
	}

	usr,err := h.userRepo.Find(reqLogin.Email, reqLogin.Password)
	if err != nil{
		util.SendError(w,http.StatusUnauthorized, "unauthorized")
	}
	// usr := database.Find(reqLogin.Password,reqLogin.Email)
	// if usr == nil{
	// 	http.Error(w, "invalid credential",400)
	// 	return
	// }

		cnf := config.GetConfig()

	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID, 
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})

	if err != nil {
		http.Error(w, "Interal Server Error", http.StatusInternalServerError) 
		return
	}


	util.SendData(w, accessToken, http.StatusCreated)
}