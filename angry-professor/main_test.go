package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestClassDoesNotStart(t *testing.T) {
	var (
		//numberOfStudents      int   = 4
		cancellationThreshold int   = 3
		studentTimes          []int = []int{-1, -3, 4, 2}
	)

	startClass := shouldStartClassGiven(cancellationThreshold, studentTimes)

	if startClass == true {
		t.Fatal("Not enugh students arrived on time")
	}
}

func TestClassStarts(t *testing.T) {
	var (
		//numberOfStudents      int   = 4
		cancellationThreshold int   = 2
		studentTimes          []int = []int{0, -1, 2, 1}
	)

	startClass := shouldStartClassGiven(cancellationThreshold, studentTimes)

	if startClass == false {
		t.Fatal("Too many students did not arrive on time")
	}
}

func TestCorrectInput(t *testing.T) {
	inputs := readInputs(strings.NewReader("4 3\n-1 -3 4 2"))

	if inputs.cancellationThreshold != 3 {
		t.Fatal("Incorrect cancellation threshold")
	}

	if reflect.DeepEqual(inputs.studentTimes, []int{-1, -3, 4, 2}) == false {
		t.Fatal("Incorrect student times")
	}
}

func TestCorrectInput2(t *testing.T) {
	inputs := readInputs(strings.NewReader("4 2\n0 -1 2 1"))

	if inputs.cancellationThreshold != 2 {
		t.Fatal("Incorrect cancellation threshold")
	}

	if reflect.DeepEqual(inputs.studentTimes, []int{0, -1, 2, 1}) == false {
		t.Fatal("Incorrect student times")
	}
}
