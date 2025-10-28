package main

import "fmt"

func main() {
	msg := sayHello("World")
	fmt.Println(msg)
}

func sayHello(name string) string {
	return "Hello, " + name + "! <3"
}
