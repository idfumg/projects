package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	Message string `json:"message"`
	Author  string `json:"-"`          // do not output this
	Date    string `json:",omitempty"` // do not output if empty
	Id      int    `json:"id,string"`  // convert to a string and rename
}

type helloWorldRequest struct {
	Name string `json:"name"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)
	catHandler := http.FileServer(http.Dir("./images"))
	http.Handle("/cat/", http.StripPrefix("/cat/", catHandler))

	log.Printf("Server starting on port: %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message: "HelloWorld"}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
