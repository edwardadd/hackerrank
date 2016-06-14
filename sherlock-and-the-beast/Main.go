package main

import "fmt"

func memset(s []byte, c byte, offset int, n int) []byte {
	for i := offset; i < offset+n; i++ {
		s[i] = c
	}
	return s
}

func FindDecentNumber(n int) string {

	if n < 3 {
		return "-1"
	}

	remainder5s := n % 3
	remainder3s := n % 5

	total5s := n / 3
	total3s := n / 5

	var result []byte = make([]byte, n)

	// how many 5s can we fit first
	if remainder5s == 0 && total5s > 0 {
		result = memset(result, '5', 0, total5s*3)
		return string(result)
	}

	if remainder3s == 0 && total3s > 0 {
		result = memset(result, '3', 0, total3s*5)
		return string(result)
	}

	// see if we can fill the rest of the spaces with 3s
	// if 5s can't fit see whether 3s can

	for i := total5s * 3; i >= 0; i -= 3 {
		space := (n - i) / 5
		spacer := (n - i) % 5

		if spacer == 0 {
			result = memset(result, '5', 0, i)
			result = memset(result, '3', i, space*5)
			return string(result)
		}
	}

	// otherwise quit

	return "-1"
}

func main() {
	var no int
	fmt.Scanf("%d", &no)

	for i := 0; i < no; i++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(FindDecentNumber(n))
	}
}
