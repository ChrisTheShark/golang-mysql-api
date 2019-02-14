package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	user := User{}
	assert.True(t, user.IsEmpty())
}

func TestIsEmptyNamePresent(t *testing.T) {
	user := User{
		Name: "Jon",
	}
	assert.False(t, user.IsEmpty())
}

func TestIsEmptyGenderPresent(t *testing.T) {
	user := User{
		Gender: "female",
	}
	assert.False(t, user.IsEmpty())
}

func TestIsEmptyAgePresent(t *testing.T) {
	user := User{
		Age: 21,
	}
	assert.False(t, user.IsEmpty())
}

func TestIsEmptyIdPresent(t *testing.T) {
	user := User{
		ID: "324545",
	}
	assert.False(t, user.IsEmpty())
}

func TestIsEmptyAllFields(t *testing.T) {
	user := User{
		Name:   "Jen",
		Gender: "female",
		Age:    43,
		ID:     "324545",
	}
	assert.False(t, user.IsEmpty())
}
