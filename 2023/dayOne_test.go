package main

import (
	"testing"
)

func TestParseFile(t *testing.T) {
	expected := 142
	actual := ParseFile("testInput.txt")

	if actual != expected {
		t.Error("expected ", expected, "but recieved ", actual)
	}
}
