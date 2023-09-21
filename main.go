package main

import (
	"fmt"
	"h8-assignment-2/handler"
)

type SimpleMethod interface {
	Greet()
}

type Person struct{}

type MyString string

func (p Person) Greet() {
	fmt.Println("ini dari person")
}

func (m MyString) Greet() {
	fmt.Println("ini dari my string")
}

func main() {
	handler.StartApp()

	// var a SimpleMethod = Person{}

	// var b SimpleMethod = MyString("my string")

	// a.Greet()

	// b.Greet()
}
