package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Login struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"pass"`
}

var users = []Login{
	{
		ID:       "1",
		Email:    "iis@gmail.com",
		Password: "lovespiderman",
	},
	{
		ID:       "2",
		Email:    "ais@gmail.com",
		Password: "lovesyou",
	},
}

func FindAcc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetAcc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var account Login
	var isFound = false

	for _, get := range users {
		if id == get.ID {
			isFound = true
			account = get
		}
	}

	if isFound == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Account with ID " + id + " not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

func MakeAcc(w http.ResponseWriter, r *http.Request) {
	var newuser Login

	json.NewDecoder(r.Body).Decode(&newuser)

	users = append(users, newuser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func DeleteAcc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var isFound = false
	var index = 0

	for i, clear := range users {
		if id == clear.ID {
			isFound = true
			index = i
		}
	}

	if isFound == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Account with ID " + id + " can't be deleted")
		return
	}

	users = append(users[:index], users[index+1:]...)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("ID " + id + " deletion success!")
}

func EditAcc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var data Login
	var isFound = false

	json.NewDecoder(r.Body).Decode(&data)

	for i, found := range users {
		if id == found.ID {
			isFound = true
			users[i] = data
		}
	}

	if isFound == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Failed to update.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
