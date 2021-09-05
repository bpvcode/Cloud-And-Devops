package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Email       string       `json:"email"`
	Age         int          `json:"age"`
	UserProfile *UserProfile `json:"userprofile"`
}

type UserProfile struct {
	Standard bool `json:"standard"`
	Premium  bool `json:"premium"`
	Gold     bool `json:"gold"`
}

var users []User

const port string = ":8000"

func main() {

	// Init Router
	r := mux.NewRouter()

	// Mock Database
	users = append(users, User{Id: 1, Name: "Bruno", Email: "bruno@gmail.com", Age: 27, UserProfile: &UserProfile{Standard: true, Premium: false, Gold: false}})
	users = append(users, User{Id: 1, Name: "John", Email: "john@gmail.com", Age: 40, UserProfile: &UserProfile{Standard: true, Premium: true, Gold: false}})
	users = append(users, User{Id: 1, Name: "Doe", Email: "doe@gmail.com", Age: 52, UserProfile: &UserProfile{Standard: true, Premium: true, Gold: true}})
	users = append(users, User{Id: 1, Name: "Hydra", Email: "hydra@gmail.com", Age: 65, UserProfile: &UserProfile{Standard: true, Premium: true, Gold: false}})
	users = append(users, User{Id: 1, Name: "Lewis", Email: "lewis@gmail.com", Age: 20, UserProfile: &UserProfile{Standard: true, Premium: false, Gold: false}})

	// Routing
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users/premium", getUsersPremium).Methods("GET")
	r.HandleFunc("/api/users/gold", getUsersGold).Methods("GET")

	log.Printf("Starting server on: `http://localhost%v` \n", port)

	// Server Listen and Serve
	log.Fatal(http.ListenAndServe(port, r))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUsersPremium(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var usersPremium []User

	for _, user := range users {
		if user.UserProfile.Premium {
			usersPremium = append(usersPremium, user)
		}
	}

	json.NewEncoder(w).Encode(usersPremium)
}

func getUsersGold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var usersGold []User

	for _, user := range users {
		if user.UserProfile.Gold {
			usersGold = append(usersGold, user)
		}
	}

	json.NewEncoder(w).Encode(usersGold)
}
