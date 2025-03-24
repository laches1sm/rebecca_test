package main

import (
	"fmt"
)


func sayHello(greeting string) string{
	return greeting
}

func main() {
	fmt.Println(sayHello("Hello, World!"))
}
