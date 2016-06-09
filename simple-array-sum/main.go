package main

import "fmt"

func main() {
	var total int = 0
	var count int = 0
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		var num int
		fmt.Scan(&num)

		total += num
	}

	fmt.Println(total)
}
