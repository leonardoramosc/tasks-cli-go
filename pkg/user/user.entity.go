package user

import (
	"errors"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Age      int8   `json:"age"`
	Password string `json:"password"`
}

func (u *User) validate() error {
	if e := validateUsername(u.Username); e != nil {
		return e
	}
	if e := validateAge(u.Age); e != nil {
		return e
	}
	if e := validatePassword(u.Password); e != nil {
		return e
	}
	return nil
}

func CreateUser(username string, age int8, password string) (*User, error) {
	newUser := &User{ Username: username, Age: age, Password: password }
	if e := newUser.validate(); e != nil {
		return nil, e
	}

	return newUser, nil
}

func validatePassword(password string) error {
	const passwordMinLength = 5
	if len(password) < passwordMinLength {
		return errors.New(
			fmt.Sprintf("Contraseña debe contener mínimo %d caracteres", passwordMinLength),
		)
	}
	return nil
}

func validateUsername(username string) error {
	const usernameMinLength int = 3
	if len(username) < usernameMinLength {
		return errors.New(
			fmt.Sprintf("El nombre de usuario debe contener mínimo %d caracteres", usernameMinLength),
		)
	}
	return nil
}

func validateAge(age int8) error {
	const maxAge int8 = 110
	
	if age < 0 || age > maxAge {
		return errors.New(
			fmt.Sprintf("La edad del usuario debe ser mayor a 0 y menor que %d", maxAge),
		)
	}

	return nil
}
