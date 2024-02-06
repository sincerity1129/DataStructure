package bubblesort_test

import (
	"math/rand"
	"testing"
)

func SetBubble(arr []int) {
	arrLength := len(arr)

	for i := 0; i < arrLength; i++ {
		for j := 0; j < arrLength-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
}

func TestBubble(t *testing.T) {
	listLength := 100
	originalList := make([]int, listLength)
	for i := 0; i < len(originalList); i++ {
		randNum := rand.Intn(100)
		originalList[i] = randNum
	}
	SetBubble(originalList)
	t.Log(originalList)
}
