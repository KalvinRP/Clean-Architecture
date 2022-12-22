package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/user", h.FindAcc).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetAcc).Methods("GET")
	r.HandleFunc("/user", h.MakeAcc).Methods("POST")
	r.HandleFunc("/user/{id}", h.EditAcc).Methods("PATCH")
	r.HandleFunc("/user", h.DeleteAcc).Methods("DELETE")
}
