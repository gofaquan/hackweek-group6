package model

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Status int
	Msg    string
	Data   interface{}
}

func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}
