package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func handleRootPath(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return;
	}

	w.Header().Set("Content-Type", "text/plain")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Person API! Use /persons to get the list of persons."))
}


func handleGetPersons(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return;
	}

	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	encoded := json.NewEncoder(w)
	encoded.Encode(people)
}


 type Person struct {
	Name string `json:"name`
	Age  int  `json:"age"`
 }

 var people = []Person{}

 func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleRootPath)

	mux.HandleFunc("/persons", handleGetPersons)

	fmt.Println("Starting server at :8080")

	err:= http.ListenAndServe(":8080", mux)

	if err!= nil{
		fmt.Println(err)
	}


	}


	func init(){
		people = append(people, Person{Name: "Alice", Age: 30})
		people = append(people, Person{Name: "Bob", Age: 25})
		people = append(people, Person{Name: "Charlie", Age: 35})
		people = append(people, Person{Name: "Diana", Age: 28})
		people = append(people, Person{Name: "Ethan", Age: 32})
		people = append(people, Person{Name: "Fiona", Age: 27})
		people = append(people, Person{Name: "George", Age: 29})
		people = append(people, Person{Name: "Hannah", Age: 31})
		people = append(people, Person{Name: "Ian", Age: 26})
		people = append(people, Person{Name: "Jane", Age: 33})
	}