package repository

import (
	"errors"
	"strconv"

	"github.com/ChrisTheShark/golang-mysql-api/models"
	"github.com/ChrisTheShark/golang-mysql-api/repository"
)

// MockUserRepository houses logic to retrieve users from a mock repository
type MockUserRepository struct{}

// NewMockUserRepository convience function to create a MockUserRepository
func NewMockUserRepository() repository.UserRepository {
	return &MockUserRepository{}
}

var users = map[string]models.User{
	"1": models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    44,
		ID:     "1",
	},
}

// GetAll get all users from the repository
func (r MockUserRepository) GetAll() ([]models.User, error) {
	userList := []models.User{}
	for _, user := range users {
		userList = append(userList, user)
	}
	return userList, nil
}

// GetByID get a user by string identifier
func (r MockUserRepository) GetByID(id string) (*models.User, error) {
	user, ok := users[id]
	if !ok {
		return nil, models.UserNotFoundError{
			Message: "not found",
		}
	}
	return &user, nil
}

// Create a User to the repository
func (r MockUserRepository) Create(user models.User) (string, error) {
	user.ID = strconv.Itoa(len(users) + 1)
	users[user.ID] = user
	return user.ID, nil
}

// Delete a User from the repository
func (r MockUserRepository) Delete(user models.User) error {
	delete(users, user.ID)
	return nil
}

// MockErroringUserRepository returns errors for all operations.
type MockErroringUserRepository struct{}

// NewMockErroringUserRepository convience function to create a MockErroringUserRepository
func NewMockErroringUserRepository() repository.UserRepository {
	return &MockErroringUserRepository{}
}

// GetAll get all users from the repository
func (r MockErroringUserRepository) GetAll() ([]models.User, error) {
	return nil, errors.New("blamo")
}

// GetByID get a user by string identifier
func (r MockErroringUserRepository) GetByID(id string) (*models.User, error) {
	return nil, errors.New("blamo")
}

// Create a User to the repository
func (r MockErroringUserRepository) Create(user models.User) (string, error) {
	return "", errors.New("blamo")
}

// Delete a User from the repository
func (r MockErroringUserRepository) Delete(user models.User) error {
	return errors.New("blamo")
}
