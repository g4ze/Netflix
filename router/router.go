package router

import (
	"controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/{}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/{}", controller.DeleteOneMovie).Methods("DELETE")
	return router

}
