package main

import (
	"fmt"
	"testing"
)

func TestHourglassSum(t *testing.T) {
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}
	got := HourglassSum(arr)
	var want int32 = 19

	if got != want {
		t.Errorf("expected '%d' but got '%d'", want, got)
	}
}

func ExampleHourglassSum() {
	arr := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 9, 2, -4, -4, 0},
		{0, 0, 0, -2, 0, 0}}
	sum := HourglassSum(arr)
	fmt.Println(sum)
	// Output: 13
}
