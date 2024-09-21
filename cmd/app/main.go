package main

import (
	"GOproject/internal/database"
	"GOproject/internal/handlers"
	"GOproject/internal/messagesService"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&messagesService.Message{})

	repo := messagesService.NewMessageRepository(database.DB)
	service := messagesService.NewService(repo)

	handler := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/get", handler.GetMessagesHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostMessageHandler).Methods("POST")
	router.HandleFunc("/api/delete", handler.DeleteMessageHandler).Methods("DELETE")
	router.HandleFunc("/api/patch", handler.PutMessageHandler).Methods("PATCH")
	http.ListenAndServe(":8080", router)
}
