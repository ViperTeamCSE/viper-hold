package main

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	expected := "Hello George"
	actual := sayHello("George")
	if actual != expected {
		t.Errorf("Test failed. Expected: '%s',  Got: '%s'", expected, actual)
	}
}
