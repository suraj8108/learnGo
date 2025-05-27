package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
type Movie struct {
	ID       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
}

var AllMovies = []Movie{}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AllMovies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var movieDetails Movie
	var found bool
	for _, movie := range AllMovies {
		if movie.ID == params["id"] {
			movieDetails = movie
			found = true
		}
	}

	if !found {
		json.NewEncoder(w).Encode(`{"message": "No such movie exists"}`)

	} else {
		json.NewEncoder(w).Encode(movieDetails)
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieDetails Movie
	json.NewDecoder(r.Body).Decode(&movieDetails)
	movieDetails.ID = strconv.Itoa(rand.Intn(10000000))
	AllMovies = append(AllMovies, movieDetails)

	json.NewEncoder(w).Encode(movieDetails)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movieDetails Movie
	json.NewDecoder(r.Body).Decode(&movieDetails)

	params := mux.Vars(r)

	for index, movieInfo := range AllMovies {
		if movieInfo.ID == params["id"] {
			AllMovies = append(AllMovies[:index], AllMovies[index+1:]...)
		}
	}

	movieDetails.ID = strconv.Itoa(rand.Intn(10000000))
	AllMovies = append(AllMovies, movieDetails)

	json.NewEncoder(w).Encode(`{"message" : "Decoded successfully"}`)

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var id = params["id"]
	var index = -1

	for movIndex, movie := range AllMovies {
		if movie.ID == id {
			index = movIndex
			break
		}
	}

	if index == -1 {
		response := map[string]interface{}{
			"message": "Invalue Id",
			"status":  http.StatusBadRequest,
		}
		json.NewEncoder(w).Encode(response)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		AllMovies = append(AllMovies[:index], AllMovies[:index]...)
		json.NewEncoder(w).Encode(`{"message" : "Deleted movie Successfully"}`)
	}
}
