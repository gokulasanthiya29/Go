package main

import (
	"fmt"
)

/*
func calc(x, y int) (output1, output2 int) {
	output1 = x + y
	output2 = x - y
	return
} */

func calc(x, y int) (int, int) {
	var output1 = x + y
	var output2 = x - y
	return output1, output2
}

func main() {
	var a, b int
	a, b = 6, 2

	result1, result2 := calc(a, b)

	fmt.Print(result1, result2)

}
