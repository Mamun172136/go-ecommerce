package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r*http.Request){
		start :=  time.Now()
		next.ServeHTTP(w,r)

		
		diff := time.Since(start)
		fmt.Println(r.Method, r.URL.Path, diff)
	})
}