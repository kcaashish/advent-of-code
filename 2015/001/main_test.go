package main

import "testing"

func TestGetCurrentFloor(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want int
	}{
		{"a", "(())", 0},
		{"b", "()()", 0},
		{"c", "(((", 3},
		{"d", "(()(()(", 3},
		{"e", "))(((((", 3},
		{"f", "())", -1},
		{"g", "))(", -1},
		{"h", ")))", -3},
		{"i", ")())())", -3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := getCurrentFloor(0, test.in)
			if out != test.want {
				t.Errorf("got %d, want %d", out, test.want)
			}
		})
	}
}

func TestGetBasementFirstEntry(t *testing.T) {
	var tests = []struct {
		name string
		in   string
		want int
	}{
		{"a", "(()))", 5},
		{"b", "())()", 3},
		{"c", "())", 3},
		{"d", "))(", 1},
		{"e", ")))", 1},
		{"f", "()(()))())", 7},
		{"g", "()", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := getFloorEntry(0, test.in, -1)
			if out != test.want {
				t.Errorf("got %d, want %d", out, test.want)
			}
		})
	}
}
