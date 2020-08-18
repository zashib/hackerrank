package main

// https://www.hackerrank.com/challenges/arrays-ds/problem
func ReverseArray(numbers []int32) []int32 {
	newNumbers := make([]int32, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}
