package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
func GetConnection(cnf *config.DBConfig)string{

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",cnf.User,cnf.Password,cnf.Host,cnf.Port,cnf.Name)
	if !cnf.EnableSSLMODE{
		connString += " sslmode=disable"
	}
	return connString
	// return "user=postgres password=1234 host=localhost port=5432 dbname=ecommerce sslmode=disable"
}

func NewConnection(cnf *config.DBConfig)(*sqlx.DB, error){

	dbsource := GetConnection(cnf)

	dbcon,err:= sqlx.Connect("postgres",dbsource)
		if err!= nil{
			fmt.Println(err)
			return nil,err
	}
	return dbcon, nil
}