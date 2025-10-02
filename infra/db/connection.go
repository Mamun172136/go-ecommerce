package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)
func GetConnection()string{

	return "user=postgres password=1234 host=localhost port=5432 dbname=ecommerce"
}

func NewConnection()(*sqlx.DB, error){

	dbsource := GetConnection()

	dbcon,err:= sqlx.Connect("postgres",dbsource)
		if err!= nil{
			fmt.Println(err)
			return nil,err
	}
	return dbcon, nil
}