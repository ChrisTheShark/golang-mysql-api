package models

// User type represents a person using the system.
type User struct {
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`
	Age    int    `json:"age" bson:"age"`
	ID     string `json:"id" bson:"_id"`
}

// IsEmpty returns a boolean value representing if the object is empty.
func (u User) IsEmpty() bool {
	return u.Name == "" && u.Gender == "" && u.Age == 0 && u.ID == ""
}

// UserNotFoundError identifies when a user is not found
type UserNotFoundError struct {
	Message string
}

func (u UserNotFoundError) Error() string {
	return u.Message
}
