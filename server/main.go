package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


 type Person struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
 }
  var people = []Person{}

func handleRootPath(w http.ResponseWriter, r *http.Request){
	handleHeaders(w)
	w.Write([]byte("Welcome to the Person API! Use /persons to get the list of persons."))
}


func handleGetPersons(w http.ResponseWriter, r *http.Request){

	handleHeaders(w)

	handleEncoder(w, people)
}

func handleCreatePerson(w http.ResponseWriter, r *http.Request){
		handleHeaders(w)

	var newPerson Person

	err:= handleDecoder(r, &newPerson)
	if err!= nil{
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	people = append(people, newPerson)
	w.WriteHeader(http.StatusCreated)
	handleEncoder(w, newPerson)

}

 func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(handleRootPath))

	mux.Handle("GET /persons", http.HandlerFunc(handleGetPersons))
	mux.Handle("POST /persons/create", http.HandlerFunc(handleCreatePerson))

	fmt.Println("Starting server at :8080")

	err:= http.ListenAndServe(":8080", mux)

	if err!= nil{
		fmt.Println(err)
	}


	}

// Predefined list of people
	var peoples = []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
		{Name: "Diana", Age: 28},
		{Name: "Ethan", Age: 32},
		{Name: "Fiona", Age: 27},
		{Name: "George", Age: 29},
		{Name: "Hannah", Age: 31},
		{Name: "Ian", Age: 26},
		{Name: "Jane", Age: 33},
	}

	// Initialize the people slice with predefined data
	func init(){
		people = peoples
	} 



	// Helper functions for handling headers, encoding, and decoding

	func handleHeaders(w http.ResponseWriter){
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	}

	func handleEncoder(w http.ResponseWriter, data interface{}){
		encoded := json.NewEncoder(w)
		encoded.Encode(data)
	}

	func handleDecoder(r *http.Request, data interface{}) error{
		err:= json.NewDecoder(r.Body).Decode(data)
		if err!= nil{
			return err
		}
		return nil
	}