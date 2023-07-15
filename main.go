package main

import (
	"log"
	"net/http"

	"github.com/muhardiansyah15/go-rest-api/config"
	"github.com/muhardiansyah15/go-rest-api/controllers"
	"github.com/muhardiansyah15/go-rest-api/middlewares"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()
	r := mux.NewRouter()
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/logout", controllers.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", controllers.Index).Methods("GET")
	api.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
