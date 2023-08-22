package app

import (
	"github.com/Firgisotya/go-commerce/app/controllers"
	"github.com/gorilla/mux"
)

func (server *Server) initiaLizeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}