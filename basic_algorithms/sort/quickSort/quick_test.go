package quicksort_test

import (
	"math/rand"
	"testing"
)

func SetQuick(arr []int, standard int) []int {
	if len(arr) <= 1 {
		return arr
	}

	if standard != 0 || len(arr) > standard {
		standard = 0
	}

	pivot := arr[standard]
	arrSorted := map[string][]int{"left": nil, "middle": nil, "right": nil}
	for _, x := range arr {
		switch {
		case x < pivot:
			arrSorted["left"] = append(arrSorted["left"], x)
		case x == pivot:
			arrSorted["middle"] = append(arrSorted["middle"], x)
		case x > pivot:
			arrSorted["right"] = append(arrSorted["right"], x)
		}
	}

	return append(append(SetQuick(arrSorted["left"], standard), arrSorted["middle"]...),
		SetQuick(arrSorted["right"], standard)...)
}

func TestQuickSort(t *testing.T) {
	original_list := make([]int, 100)

	for i := 0; i < len(original_list); i++ {
		randomNum := rand.Intn(100)
		original_list[i] = randomNum
	}
	result := SetQuick(original_list, 10)
	t.Log(result)
}
