package controller

import (
	"encoding/json"
	"net/http"

	"github.com/JuliusSi/movie-db/movies/entity"
)

func CreateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
}

func GetMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	var Movies []entity.Movie
	Movies = append(Movies, entity.Movie{"Rajeev", "Singh", "Test"})

	json.NewEncoder(res).Encode(Movies)
}
