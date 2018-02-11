package main

/**
Go version:1.9.3

@Feb 2018
@by Ashwin Rana
*/

import (
	"testing"
)

func Test_most_Sold(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			most_Sold()
		})
	}
}

func Test_total_Spend(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			total_Spend(tt.args.email)
		})
	}
}

func Test_most_Loyal(t *testing.T) {
	tests := []struct {
		name string
	}{
	// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			most_Loyal()
		})
	}
}
