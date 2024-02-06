package selectsort_test

import (
	"math/rand"
	"testing"
)

func setSelect(arr []int) {
	arrLength := len(arr)
	for i := 0; i < arrLength; i++ {
		minIndex := i

		for j := i + 1; j < arrLength; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		tmp := arr[minIndex]
		arr[minIndex] = arr[i]
		arr[i] = tmp
	}
}

func TestSelect(t *testing.T) {
	listLength := 100
	originalList := make([]int, listLength)
	for i := 0; i < len(originalList); i++ {
		randNum := rand.Intn(100)
		originalList[i] = randNum
	}
	setSelect(originalList)
	t.Log(originalList)
}
