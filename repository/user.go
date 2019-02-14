package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/ChrisTheShark/golang-mysql-api/models"
)

// UserRepository inteface describes rrepository operations on Users
type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(string) (*models.User, error)
	Create(models.User) (string, error)
	Delete(models.User) error
}

// UserRepositoryImpl houses logic to retrieve users from a mongo repository
type UserRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository convience function to create a UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{db}
}

// GetAll get all users from the repository
func (r UserRepositoryImpl) GetAll() ([]models.User, error) {
	users := []models.User{}

	rows, err := r.db.Query("select id, name, age, gender from users")
	if err != nil {
		return nil, fmt.Errorf("unable to locate users due to: %v", err)
	}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
			return nil, fmt.Errorf("unable to locate users due to: %v", err)
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByID get a user by string identifier
func (r UserRepositoryImpl) GetByID(id string) (*models.User, error) {
	var user models.User
	row := r.db.QueryRow("select id, name, age, gender from users where id = ?", id)
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Gender); err != nil {
		return nil, fmt.Errorf("unable to locate user due to: %v", err)
	}
	return &user, nil
}

// Create a User to the repository
func (r UserRepositoryImpl) Create(user models.User) (string, error) {
	result, err := r.db.Exec("insert into users (name, age, gender) values (?, ?, ?)",
		user.Name, user.Age, user.Gender)
	if err != nil {
		return "", fmt.Errorf("unable to create user due to: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("unable to create user due to: %v", err)
	}
	return strconv.FormatInt(id, 10), nil
}

// Delete a User from the repository
func (r UserRepositoryImpl) Delete(user models.User) error {
	result, err := r.db.Exec("delete from users where id = ?", user.ID)
	if err != nil {
		return fmt.Errorf("unable to delete user due to: %v", err)
	}

	re, err := result.RowsAffected()
	if err != nil || re != 1 {
		return fmt.Errorf("unable to delete user due to: %v", err)
	}
	return nil
}
