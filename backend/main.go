package main

import (
	"github.com/aryanstha/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Auth Routes
	r.HandleFunc("/api/login", handlers.login).Methods("POST")

	// License Routes
	r.Hand
	eFunc("/api/license", middleware.AuthMiddleware(handlers.createLicense)).Methods("POST")
	r.Hand
	eFunc("/api/license", middleware.AuthMiddleware(handlers.getLicenses)).Methods("GET")
	r.Hand
	eFunc("/api/license/{id}", middleware.AuthMiddleware(handlers.getLicense)).Methods("GET")
	r.Hand
	eFunc("/api/license/{id}", middleware.AuthMiddleware(handlers.updateLicense)).Methods("PUT")
	r.Hand
	eFunc("/api/license/{id}", middleware.AuthMiddleware(handlers.deleteLicense)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
