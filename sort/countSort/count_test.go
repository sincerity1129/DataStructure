package countsort_test

import (
	"math/rand"
	"testing"
)

func SetCountSort(arr []int) []int {
	maxValue := arr[0]

	for _, value := range arr {
		if maxValue < value {
			maxValue = value
		}
	}

	defaultCount := make([]int, maxValue+1)
	for _, value := range arr {
		defaultCount[value] += 1
	}

	for i := 1; i < len(defaultCount); i++ {
		defaultCount[i] += defaultCount[i-1]
	}

	sortedCount := make([]int, len(arr))
	for _, value := range arr {
		sortedCount[defaultCount[value]-1] = value
		defaultCount[value] -= 1
	}
	return sortedCount
}

func TestCount(t *testing.T) {
	original_list := make([]int, 10)

	for i := 0; i < len(original_list); i++ {
		randNum := rand.Intn(100)
		original_list[i] = randNum
	}

	sortedList := SetCountSort(original_list)
	t.Log(original_list)
	t.Log(sortedList)
}
