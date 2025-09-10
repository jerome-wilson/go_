package main

import "testing"

func TestStruct(t *testing.T) {
	// arrange
	l, r := 2, 3
	expect := 5

	// act
	got := Add(l, r)

	// assert
	if expect != got {
		t.Errorf("Failed to add %v and %v . Got %v, expected %v\n", l, r, got, expect)
	}
}