package main

import (
	"fmt"
	"io"
	"os"
)

type Input struct {
	cancellationThreshold int
	studentTimes          []int
}

func readInputs(reader io.Reader) Input {
	var input Input
	var numberOfStudentTimes int = 0

	if _, err := fmt.Fscan(reader, &numberOfStudentTimes, &input.cancellationThreshold); err != nil {
		return Input{0, []int{}}
	}

	input.studentTimes = make([]int, numberOfStudentTimes)
	for i := 0; i < numberOfStudentTimes; i++ {
		var time int = 0
		if _, err := fmt.Fscan(reader, &time); err != nil {
			return Input{0, []int{}}
		}
		input.studentTimes[i] = time
	}

	return input
}

func shouldStartClassGiven(t int, times []int) bool {
	var onTimeCount int = 0
	for _, time := range times {
		if time <= 0 {
			onTimeCount++
		}
	}

	if onTimeCount >= t {
		return true
	}

	return false
}

func main() {
	var numberOfTestCases int = 0
	fmt.Fscan(os.Stdin, &numberOfTestCases)

	for i := 0; i < numberOfTestCases; i++ {
		input := readInputs(os.Stdin)
		cancel := !shouldStartClassGiven(input.cancellationThreshold, input.studentTimes)
		if cancel {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
