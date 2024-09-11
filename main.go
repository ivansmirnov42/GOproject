package main

import (
	"GOproject/db"
	"GOproject/orm"
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
	w.WriteHeader(http.StatusOK)

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	message = reqBody.Message
	db.DB.Create(&orm.Message{Text: message})

}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var message []orm.Message
	db.DB.Find(&message)
	json.NewEncoder(w).Encode(message)
}

func PatchMessage(w http.ResponseWriter, r *http.Request) {
	var id string
	var new_message string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	new_message = reqBody.Message
	id = reqBody.Id
	db.DB.Model(&orm.Message{}).Where("id = ?", id).Update("text", new_message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	var id string
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var reqBody requestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println(err)
	}
	id = reqBody.Id
	db.DB.Where("id = ?", id).Delete(&orm.Message{})
}

func main() {
	db.InitDB()
	err := db.DB.AutoMigrate(&orm.Message{})

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
