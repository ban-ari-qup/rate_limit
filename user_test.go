package main

import (
	"errors"
	"testing"
	"time"
)

func TestCanMakeReq(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want bool
	}{
		{"zero value", User{}, true},
		{"under limit", User{name: "Ali", limit: 3, Requests: []time.Time{time.Now(), time.Now()}}, true},
		{"equal to limit", User{name: "Asem", limit: 3, Requests: []time.Time{time.Now(), time.Now(), time.Now()}}, false},
		{"more than limit", User{name: "Amina", limit: 3, Requests: []time.Time{time.Now(), time.Now(), time.Now(), time.Now()}}, false},
		{"one with expired date", User{name: "Abu", limit: 3, Requests: []time.Time{time.Now().Add(-61 * time.Second), time.Now().Add(-60 * time.Second), time.Now().Add(-59 * time.Second)}}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.canMakeReq()
			if got != tt.want {
				t.Errorf("%+v\ncanMakeReq() = %t; want %t", tt.u, got, tt.want)
			}
		})
	}

	// u := &User{}
	// got := u.canMakeReq()
	// want := true
	// if got != want {
	// 	t.Errorf("canMakeReq() = %t; want %t", got, want)
	// }
}

func TestAddRequest(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want error
	}{
		{"We can add", User{name: "Ali", limit: 3, Requests: []time.Time{time.Now()}}, nil},
		{"We cant add", User{name: "Asem", limit: 3, Requests: []time.Time{time.Now(), time.Now(), time.Now()}}, errors.New("u get ur limit")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.addRequest()
			if got != tt.want {
				t.Errorf("%+v\n addRequest() = %v; want %v", tt.u, got, tt.want)
			}
		})
	}
}
