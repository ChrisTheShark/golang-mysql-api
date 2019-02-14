package repository

import (
	"testing"

	"github.com/ChrisTheShark/golang-mysql-api/models"
	"github.com/stretchr/testify/assert"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to create mock DB object: %v", err)
	}
	defer db.Close()

	expectedUsers := []models.User{
		models.User{
			ID:     "1",
			Name:   "James Bond",
			Age:    43,
			Gender: "male",
		},
	}

	rows := sqlmock.NewRows([]string{"id", "name", "age", "gender"}).
		AddRow(1, expectedUsers[0].Name, expectedUsers[0].Age, expectedUsers[0].Gender)
	mock.ExpectQuery("select (.+) from users").
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	users, err := ur.GetAll()
	if err != nil {
		t.Fatalf("unable to execute GetAll in TestGetAll due to: %v", err)
	}

	assert.Equal(t, expectedUsers, users)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to create mock DB object: %v", err)
	}
	defer db.Close()

	expectedUser := models.User{
		ID:     "1",
		Name:   "James Bond",
		Age:    43,
		Gender: "male",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "age", "gender"}).
		AddRow(1, expectedUser.Name, expectedUser.Age, expectedUser.Gender)
	mock.ExpectQuery("select (.+) from users").
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	user, err := ur.GetByID("1")
	if err != nil {
		t.Fatalf("unable to execute GetByID in TestGetByID due to: %v", err)
	}

	assert.NotEmpty(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Name, user.Name)
	assert.Equal(t, expectedUser.Age, user.Age)
	assert.Equal(t, expectedUser.Gender, user.Gender)
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to create mock DB object: %v", err)
	}
	defer db.Close()

	expectedUser := models.User{
		Name:   "James Bond",
		Age:    43,
		Gender: "male",
	}

	mock.ExpectExec("insert into users").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepository(db)
	id, err := ur.Create(expectedUser)
	if err != nil {
		t.Fatalf("unable to execute GetByID in TestGetByID due to: %v", err)
	}

	assert.Equal(t, "1", id)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to create mock DB object: %v", err)
	}
	defer db.Close()

	expectedUser := models.User{
		ID:     "1",
		Name:   "James Bond",
		Age:    43,
		Gender: "male",
	}

	mock.ExpectExec("delete from users").
		WillReturnResult(sqlmock.NewResult(1, 1))

	ur := NewUserRepository(db)
	err = ur.Delete(expectedUser)
	if err != nil {
		t.Fatalf("unable to execute GetByID in TestGetByID due to: %v", err)
	}
}
