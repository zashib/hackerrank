package main

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	numbers := []int32{1, 4, 3, 2}
	got := ReverseArray(numbers)
	want := []int32{2, 3, 4, 1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func BenchmarkReverseArray(b *testing.B) {
	numbers := []int32{1, 4, 3, 2}
	for i := 0; i < b.N; i++ {
		ReverseArray(numbers)
	}
}
