package main

import "fmt"

// variable scope in "go" -> package level scope and function level scope
var a int = 9
var b int = 10

func demo() {
	a := 6
	fmt.Println("demo", a, b)
}

func main() {
	demo()
	fmt.Println("main", a)
}
