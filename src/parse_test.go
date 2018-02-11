package main

/**
Go version:1.9.3

@Feb 2018
@by Ashwin Rana
*/

import (
	"reflect"
	"testing"
)

func Test_getUsers(t *testing.T) {
	tests := []struct {
		name string
		want UserData
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPurchases(t *testing.T) {
	tests := []struct {
		name string
		want PurchaseData
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPurchases(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPurchases() = %v, want %v", got, tt.want)
			}
		})
	}
}
