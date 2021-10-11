package main

import (
	"net/http"

	moviesRoutes "github.com/JuliusSi/movie-db/movies/router"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	moviesRoutes.Routes(router)

	http.ListenAndServe(":8080", router)
}
