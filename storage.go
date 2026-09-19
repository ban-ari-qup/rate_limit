package main

import (
	"math/rand"
	"time"
)

var Users = make(map[string]*User)

func CheckIfExists(apiKey string) bool {
	_, ok := Users[apiKey]
	return ok
}

func registerUser(name string) string {
	apiKey := apiGenerater()
	Users[apiKey] = createUser(name)

	return apiKey
}

func createUser(name string) *User {
	limit := generateLimit()
	var u User = User{name, limit, []time.Time{}}
	return &u
}

func generateLimit() int {
	limit := rand.Intn(60) + 30
	return limit
}

func apiGenerater() string {
	letters := "abcdefghijklmnopqrstuvwxyz1234567890"
	l := make([]byte, 10)
	for i := range l {
		l[i] = letters[rand.Intn(len(letters))]
	}
	return string(l)
}
