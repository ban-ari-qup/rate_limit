package main

import (
	"errors"
	"time"
)

type User struct {
	name     string
	limit    int
	Requests []time.Time
}

func (u *User) updateReq() []time.Time {
	return u.Requests[1:]
}

func (u *User) canMakeReq() bool {
	if len(u.Requests) == 0 {
		return true
	}
	if time.Since(u.Requests[0]) > time.Second*60 {
		u.Requests = u.updateReq()
		return u.canMakeReq()
	} else {
		if len(u.Requests) >= u.limit {
			return false
		} else {
			return true
		}
	}
}

func (u *User) addRequest() error {
	temp := u.canMakeReq()
	if temp {
		u.Requests = append(u.Requests, time.Now())
		return nil
	} else {
		return errors.New("u get ur limit")
	}
}
