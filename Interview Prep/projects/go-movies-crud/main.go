package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/suraj8108/movieSystem/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", handler.GetMovies).Methods("GET")
	r.HandleFunc("/movie/{id}", handler.GetMovie).Methods("GET")
	r.HandleFunc("/movie", handler.CreateMovie).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", handler.DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server on 8000")
	fmt.Println(http.ListenAndServe(":8000", r))

}
