package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/ChrisTheShark/golang-mysql-api/models"
	mocks "github.com/ChrisTheShark/golang-mysql-api/repository/mocks"
	"github.com/julienschmidt/httprouter"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUsers(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}

	uc := NewUserController(mocks.NewMockUserRepository())
	uc.GetUsers(w, r, p)
	resp := w.Result()

	bs, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, "[{\"name\":\"James Bond\",\"gender\":\"male\",\"age\":44,\"id\":\"1\"}]\n", string(bs))
}

func TestGetAllUsersNegativePath(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}

	uc := NewUserController(mocks.NewMockErroringUserRepository())
	uc.GetUsers(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, "503 Service Unavailable", resp.Status)
}

func TestGetUserByID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "1",
	})

	uc := NewUserController(mocks.NewMockUserRepository())
	uc.GetUserByID(w, r, p)
	resp := w.Result()

	bs, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	assert.Equal(t, "{\"name\":\"James Bond\",\"gender\":\"male\",\"age\":44,\"id\":\"1\"}\n", string(bs))
}

func TestGetUserByIDNotFound(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "99",
	})

	uc := NewUserController(mocks.NewMockUserRepository())
	uc.GetUserByID(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "404 Not Found", resp.Status)
}

func TestGetUserByIDNegativePath(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "99",
	})

	uc := NewUserController(mocks.NewMockErroringUserRepository())
	uc.GetUserByID(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, "503 Service Unavailable", resp.Status)
}

func TestAddUser(t *testing.T) {
	user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    44,
		ID:     "1",
	}

	bs, _ := json.Marshal(&user)
	r := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bs))
	w := httptest.NewRecorder()
	p := httprouter.Params{}

	uc := NewUserController(mocks.NewMockUserRepository())
	uc.AddUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusSeeOther, resp.StatusCode)
	assert.Equal(t, "303 See Other", resp.Status)
}

func TestAddUserBadRequest(t *testing.T) {
	otherStruct := struct {
		Sport      string
		NumPlayers int
	}{
		Sport:      "basketball",
		NumPlayers: 5,
	}

	bs, _ := json.Marshal(&otherStruct)
	r := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bs))
	w := httptest.NewRecorder()
	p := httprouter.Params{}

	uc := NewUserController(mocks.NewMockUserRepository())
	uc.AddUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "400 Bad Request", resp.Status)
}

func TestAddUserNegativePath(t *testing.T) {
	user := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    44,
		ID:     "1",
	}

	bs, _ := json.Marshal(&user)
	r := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(bs))
	w := httptest.NewRecorder()
	p := httprouter.Params{}

	uc := NewUserController(mocks.NewMockErroringUserRepository())
	uc.AddUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, "503 Service Unavailable", resp.Status)
}

func TestDeleteUser(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	uc := NewUserController(mocks.NewMockUserRepository())
	uc.DeleteUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	assert.Equal(t, "204 No Content", resp.Status)
}

func TestDeleteUserNotFound(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/users/99", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "99",
	})
	uc := NewUserController(mocks.NewMockUserRepository())
	uc.DeleteUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "404 Not Found", resp.Status)
}

func TestDeleteUserNegativePath(t *testing.T) {
	r := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()
	p := httprouter.Params{}
	p = append(p, httprouter.Param{
		Key:   "id",
		Value: "1",
	})
	uc := NewUserController(mocks.NewMockErroringUserRepository())
	uc.DeleteUser(w, r, p)
	resp := w.Result()

	assert.Equal(t, http.StatusServiceUnavailable, resp.StatusCode)
	assert.Equal(t, "503 Service Unavailable", resp.Status)
}
