package main

import "fmt"

// https://www.hackerrank.com/challenges/2d-array/problem
// return max sum of hourglass in array
func HourglassSum(arr [][]int32) int32 {
	var maxSum int32
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			arrTop := ArraySum(arr[i][j : j+3])
			arrMid := arr[i+1][j+1]
			arrLow := ArraySum(arr[i+2][j : j+3])
			tmpSum := arrTop + arrMid + arrLow
			if (i == 0) && (j == 0) {
				maxSum = tmpSum
			}
			if maxSum < tmpSum {
				maxSum = tmpSum
			}
		}

	}
	return maxSum

}

func ArraySum(a []int32) int32 {
	var sum int32 = 0
	for _, number := range a {
		sum += number
	}
	return sum
}

func main() {
	numbers := [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0}}
	fmt.Println(HourglassSum(numbers))
}
