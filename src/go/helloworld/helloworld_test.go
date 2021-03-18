package main

import (
	"os/user"
	"testing"
)

func TestSayHelloSuccess(t *testing.T) {
	expected := "Hello George"
	actual := sayHello("George")
	if actual != expected {
		t.Errorf("Test failed. Expected: '%s',  Got: '%s'", expected, actual)
	}
}

func TestSayHelloFail(t *testing.T) {
	expectedWrong := "Hello Mike"
	actual := sayHello("George")
	if actual == expectedWrong {
		t.Errorf("Test failed. Expected: '%s',  Got: '%s'", expectedWrong, actual)
	}
}

func TestGetCurrentUser(t *testing.T) {
	var userObject *user.User
	var err error

	userObject, err = user.Current()

	if err != nil {
		panic(err)
	}
	expected := userObject.Username
	actual := getCurrentUser()
	if actual != expected {
		t.Errorf("Test failed. Expected: '%s',  Got: '%s'", expected, actual)
	}
}

func TestMain(t *testing.T) {
	// Test will fail if main errors out
	main()
}
