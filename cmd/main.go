package main

import (
	"log"
	"net/http"

	"github.com/spookycoincidence/hx-user-service-demo/internal/handler"
	"github.com/spookycoincidence/hx-user-service-demo/internal/repository"
	"github.com/spookycoincidence/hx-user-service-demo/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	mockRepo := repository.NewMockUserRepository()
	userService := service.NewUserService(mockRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")

	log.Println("User service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
