package main

import (
	"testing"
)

// func TestCheckIfExists(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		apikey string
// 		result bool
// 	}{
// 		{"accepted", "", true},
// 		{"declined", "", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := CheckIfExists(tt.apikey)
// 			if got != tt.result {
// 				t.Errorf("CheckIfExists(%s) = %t", tt.apikey, tt.result)
// 			}
// 		})
// 	}
// }

func TestRegisterUser(t *testing.T) {
	got := registerUser("Ali")
	if got == "" {
		t.Errorf("dont generating api key")
	}
	if Users[got].name != "Ali" {
		t.Errorf("Wrong saving method")
	}

}
