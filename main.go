package main

import(
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello Worlddo!")
		AWS_ACCOUNT_ID="123456789099"
		
	})
}

func main(){
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}
