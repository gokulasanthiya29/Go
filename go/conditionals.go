package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now().Weekday()

	switch today {
	case today + 0:
		fmt.Println("Today is Saturday :)")
	case today + 1:
		fmt.Println("Well, it is day-after-tomorrow!")
	default:
		fmt.Println("Oh!, it is too far :(")
	}
	/*	if 2 + 4 == 6 {
			fmt.Println("Right answer")
		} else if 2 + 4 == 5 {
			fmt.Println("Oh, no!")
		} else {
			("Wrong answer")
		} */
}
