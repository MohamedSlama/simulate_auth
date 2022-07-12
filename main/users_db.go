package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func getDbUser() User {
	mohamedHash, err := bcrypt.GenerateFromPassword([]byte("mohamed password"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return User{userName: "mohamed", password: mohamedHash}
}
