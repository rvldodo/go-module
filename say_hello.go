package main

import "fmt"

func SayHello(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Println(SayHello("Rivaldo"))
}
