package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ChrisTheShark/golang-mysql-api/models"
	"github.com/ChrisTheShark/golang-mysql-api/repository"
	"github.com/julienschmidt/httprouter"
)

// UserController struct containing web related logic to operate on Users
type UserController struct {
	userRepository repository.UserRepository
}

// NewUserController is a convienince function to create a UserController
func NewUserController(r repository.UserRepository) *UserController {
	return &UserController{r}
}

// GetUsers retrieve all users
func (u UserController) GetUsers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID get a user by string identifier
func (u UserController) GetUserByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		if _, ok := err.(models.UserNotFoundError); ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// AddUser add a json encoded user
func (u UserController) AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.IsEmpty() {
		log.Println(err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	id, err := u.userRepository.Create(user)
	if err != nil {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/users/%v", id), http.StatusSeeOther)
}

// DeleteUser remove a user
func (u UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		if _, ok := err.(models.UserNotFoundError); ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	if err := u.userRepository.Delete(*user); err != nil {
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
