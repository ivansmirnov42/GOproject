package main

import (
	"GOproject/database"
	"GOproject/messagesSerivce"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type requestBody struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

var message string

func PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	message = reqBody.Message
	database.DB.Create(&messagesSerivce.Message{Text: message})
	w.WriteHeader(http.StatusOK)

}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var message []messagesSerivce.Message
	database.DB.Find(&message)
	json.NewEncoder(w).Encode(message)

}

func PatchMessage(w http.ResponseWriter, r *http.Request) {
	var id string
	var new_message string
	w.Header().Set("Content-Type", "application/json")

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	new_message = reqBody.Message
	id = reqBody.Id
	database.DB.Model(&messagesSerivce.Message{}).Where("id = ?", id).Update("text", new_message)
	w.WriteHeader(http.StatusOK)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var id string
	w.Header().Set("Content-Type", "application/json")

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	id = reqBody.Id
	database.DB.Where("id = ?", id).Delete(&messagesSerivce.Message{})
	w.WriteHeader(http.StatusOK)
}

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&messagesSerivce.Message{})

	if err != nil {
		fmt.Println(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/post", PostMessage).Methods("POST")
	router.HandleFunc("/get", GetMessage).Methods("GET")
	router.HandleFunc("/patch", PatchMessage).Methods("PATCH")
	router.HandleFunc("/delete", DeleteMessage).Methods("DELETE")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
