// Package main provides ...
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	tod := input[8:]

	hour, _ := strconv.Atoi(input[:2])
	if tod == "PM" {
		if hour < 12 {
			hour += 12
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}
	fmt.Printf("%02d%s", hour, input[2:8])
}
