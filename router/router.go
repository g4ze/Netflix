package router

import (
	"controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/{id}", controller.MarkAsWatched).Methods("PUT")
	return router

}
