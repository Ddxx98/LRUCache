package routers

import (
	"backend/controllers"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",controllers.Visible)
	router.HandleFunc("/get",controllers.GetHandler).Methods("GET")
	router.HandleFunc("/set",controllers.SetHandler).Methods("POST")
	return router
}