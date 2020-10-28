package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a person who uses the system
type User struct {
	gorm.Model
	Email     string
	Password  string
	FirstName string
	LastName  string
}

// NewUser create a new user
func NewUser(email, password, firstName, lastName string) (*User, error) {
	u := &User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}

	pwd, err := generatePassword(password)
	if err != nil {
		return nil, err
	}
	u.Password = pwd

	err = u.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return u, nil
}

// Validate user data
func (u *User) Validate() error {
	if u.Email == "" || u.FirstName == "" || u.LastName == "" || u.Password == "" {
		return ErrInvalidEntity
	}

	return nil
}

// ValidatePassword validate user password
func (u *User) ValidatePassword(p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return err
	}
	return nil
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
