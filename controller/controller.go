package controller

import (
	"encoding/json"
	"helper"
	"model"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "controlleeerrrr")
	allMovies := helper.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "controlleeerrrr")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Movies
	_ = json.NewDecoder(r.Body).Decode(&movie)
	helper.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}
func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "controlleeerrrr")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	helper.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "controlleeerrrr")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	helper.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
