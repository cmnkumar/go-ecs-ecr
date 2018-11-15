package main

import(
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World!")
		AWS_ACCESS_KEY_ID="AKIAIZTRXAHEBLAHKLAH"
		AWS_ACCOUNT_ID="123456789012"
	})
}

func main(){
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}
