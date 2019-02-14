package repository

import (
	"errors"
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

func TestGetAllQueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unable to create mock DB object: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery("select (.+) from users").
		WillReturnError(errors.New("blamo"))

	ur := NewUserRepository(db)
	users, err := ur.GetAll()

	assert.Nil(t, users)
	assert.NotNil(t, err)
	assert.Equal(t, "unable to locate users due to: blamo", err.Error())
}

func TestGetAllRowScanError(t *testing.T) {
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

	// Adding a value of type string to the rows for age, this should
	// trigger a row scan error as go attempts to set a string value into
	// an int field.
	rows := sqlmock.NewRows([]string{"id", "name", "age", "gender"}).
		AddRow(expectedUsers[0].ID, expectedUsers[0].Name, "expectedUsers[0].Age", expectedUsers[0].Gender)
	mock.ExpectQuery("select (.+) from users").
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	users, err := ur.GetAll()

	assert.Nil(t, users)
	assert.NotNil(t, err)
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

func TestGetByIDRowScanError(t *testing.T) {
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

	// Adding a value of type string to the rows for age, this should
	// trigger a row scan error as go attempts to set a string value into
	// an int field.
	rows := sqlmock.NewRows([]string{"id", "name", "age", "gender"}).
		AddRow(1, expectedUser.Name, "expectedUser.Age", expectedUser.Gender)
	mock.ExpectQuery("select (.+) from users").
		WillReturnRows(rows)

	ur := NewUserRepository(db)
	user, err := ur.GetByID("1")

	assert.Nil(t, user)
	assert.NotNil(t, err)
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

func TestCreateExecError(t *testing.T) {
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
		WillReturnError(errors.New("blamo"))

	ur := NewUserRepository(db)
	id, err := ur.Create(expectedUser)

	assert.Empty(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "unable to create user due to: blamo", err.Error())
}

func TestCreateRowsAffectedError(t *testing.T) {
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
		WillReturnResult(sqlmock.NewErrorResult(errors.New("blamo")))

	ur := NewUserRepository(db)
	id, err := ur.Create(expectedUser)

	assert.Empty(t, id)
	assert.NotNil(t, err)
	assert.Equal(t, "unable to create user due to: blamo", err.Error())
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

func TestDeleteExecError(t *testing.T) {
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
		WillReturnError(errors.New("blamo"))

	ur := NewUserRepository(db)
	err = ur.Delete(expectedUser)

	assert.NotNil(t, err)
	assert.Equal(t, "unable to delete user due to: blamo", err.Error())
}

func TestDeleteRowsAffectedError(t *testing.T) {
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
		WillReturnResult(sqlmock.NewErrorResult(errors.New("blamo")))

	ur := NewUserRepository(db)
	err = ur.Delete(expectedUser)

	assert.NotNil(t, err)
	assert.Equal(t, "unable to delete user due to: blamo", err.Error())
}
