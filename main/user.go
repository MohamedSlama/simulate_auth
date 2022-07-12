package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	userName string
	password []byte
}

func (user User) toString() {
	fmt.Printf("{ userName:%s ,password:%s }\n", user.userName, user.password)
}

func (user User) auth(loggedInUser User) bool {

	if !strings.Contains(user.userName, loggedInUser.userName) {
		fmt.Println("Incorrect userName!!")
		return false
	}

	if err := bcrypt.CompareHashAndPassword(user.password, loggedInUser.password); err != nil {
		fmt.Println("Incorrect password!!")
		return false
	}
	fmt.Println("logged in successfully!!")
	return true
}
