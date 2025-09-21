package main

import (
	"fmt"
	"net/http"
)


func handleAbout(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the about page")
}

 func main() {
	mux:= http.NewServeMux()
	mux.HandleFunc("/about", handleAbout)

	fmt.Println("Starting server at :8080")

	err:= http.ListenAndServe(":8080", mux)

	if err!= nil{
		fmt.Println(err)
	}



	}