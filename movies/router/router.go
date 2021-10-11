package router

import (
	moviesController "github.com/JuliusSi/movie-db/movies/controller"
	"github.com/gorilla/mux"
)

func Routes(route *mux.Router) {
	group := route.PathPrefix("/api/movies").Subrouter()
	group.HandleFunc("/create", moviesController.CreateMovie).Methods("POST")
	group.HandleFunc("", moviesController.GetMovies).Methods("GET")
}
