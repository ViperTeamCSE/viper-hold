package main

import (
	"fmt"
	"os/user"
)

func sayHello(userName string) string {
	output := "Hello "
	output += userName
	return output
}

func getCurrentUser() string {
	var userObject *user.User
	var err error

	userObject, err = user.Current()

	if err != nil {
		panic(err)
	}
	return userObject.Username
}

func main() {
	userName := getCurrentUser()
	hello := sayHello(userName)
	fmt.Print(hello)
}
