package main

import (
	"fmt"
	"os/user"
)

func sayHello(userName string) string {
	output := "Hello "
	output += userName
	return output
	// fmt.Printf("Hello %s\n", userName)
}

func main() {
	var userObject *user.User
	var err error

	userObject, err = user.Current()

	if err != nil {
		panic(err)
	}
	hello := sayHello(userObject.Username)
	fmt.Print(hello)
}
