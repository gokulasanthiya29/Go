package main

import "fmt"

func main() {
	a := 1
	const b int = 2
	var c = 6

	// b = 5
	var result int
	result = a + b + c

	fmt.Print(result)
}
