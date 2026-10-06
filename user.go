package main

import (
	"errors"
	"sync"
	"time"
)

type User struct {
	mu       sync.RWMutex
	name     string
	limit    int
	Requests []time.Time
}

func (u *User) canMakeReq() bool {
	if len(u.Requests) == 0 {
		return true
	}
	if time.Since(u.Requests[0]) > time.Second*60 {
		for len(u.Requests) > 0 && time.Since(u.Requests[0]) > time.Second*60 {
			u.Requests = u.Requests[1:]
		}

	}
	if len(u.Requests) >= u.limit {
		return false
	} else {
		return true
	}

}

func (u *User) addRequest() error {
	u.mu.Lock()
	defer u.mu.Unlock()
	temp := u.canMakeReq()
	if temp {
		u.Requests = append(u.Requests, time.Now())
		return nil
	} else {
		return errors.New("u get ur limit")
	}
}
