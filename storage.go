package main

import (
	"sync"
	"time"

	"github.com/google/uuid"
)

var mu sync.RWMutex

var generalLimit int = 60 //3

var Users = make(map[string]*User)

func CheckIfExists(apiKey string) bool {
	mu.RLock()
	defer mu.RUnlock()
	_, ok := Users[apiKey]
	return ok
}

func registerUser(name string) string {
	mu.Lock()
	defer mu.Unlock()
	apiKey := uuid.New().String()
	Users[apiKey] = createUser(name)

	return apiKey
}

func createUser(name string) *User {
	var u User = User{name: name, limit: generalLimit, Requests: []time.Time{}}
	return &u
}

func changeLimit(limit int, u *User) {
	mu.Lock()
	defer mu.Unlock()
	u.limit = limit
}
