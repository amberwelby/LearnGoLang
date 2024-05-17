// These files are only used for testing and are excluded from production builds

package main

import (
	"testing"
)

// A test is just a function, a Test test is prefixed with Test and a Fuzz test is
// prefixed with Fuzz. The first letter after Test has to be capitalized or Go
// considers it invalid
func TestAdd(t *testing.T) {
	// Arrange
	l, r := 1, 2
	expect := 3

	// Act
	got := Add(l, r)

	// Assert
	if expect != got {
		t.Errorf("Failed to add %v and %v. Got %v, expected %v\n", l, r, got, expect)
	}
}

// The t object helps communicate with the test runner

// To run, click "run package tests" at the top, or use the command line and run
// "go test ." If you want to run all of the tests in a module, use the command
// line and use "go test ./..." from the root of the module. 