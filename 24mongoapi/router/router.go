package router

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/tomarbhavuk/mongoapi/controller"
)

func Router() *mux.Router {
	fmt.Println("Learn building api")

	r := mux.NewRouter()

	//routing
	r.HandleFunc("/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/movies", controller.DeleteAllMovies).Methods("DELETE")
	r.HandleFunc("/movies/{id}", controller.DeleteAMovie).Methods("DELETE")

	return r
}
