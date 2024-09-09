package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
}

var message string

func PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	message = reqBody.Message
	DB.Create(&reqBody)

}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %v", message)
}

func main() {
	InitDB()

	err := DB.AutoMigrate(&Message{})

	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/post", PostMessage).Methods("POST")
	router.HandleFunc("/get", GetMessage).Methods("GET")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
