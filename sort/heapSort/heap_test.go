package heapsort_test

import (
	"math/rand"
	"testing"
)

func heapProcess(arr []int, arrLength int, idx int) {
	largest := idx

	leftNode := idx*2 + 1
	rightNode := idx*2 + 2

	if leftNode < arrLength && arr[leftNode] > arr[largest] {
		largest = leftNode
	}

	if rightNode < arrLength && arr[rightNode] > arr[largest] {
		largest = rightNode
	}

	if largest != idx {
		tmp := arr[largest]
		arr[largest] = arr[idx]
		arr[idx] = tmp

		heapProcess(arr, arrLength, largest)
	}
}

func setHeap(arr []int) {
	arrLength := len(arr)

	for i := (arrLength / 2) - 1; 0 <= i; i-- {
		heapProcess(arr, arrLength, i)
	}

	for i := (arrLength - 1); 0 < i; i-- {
		tmp := arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		heapProcess(arr, i, 0)
	}
}

func TestHeap(t *testing.T) {
	listLength := 100
	originalList := make([]int, listLength)
	for i := 0; i < len(originalList); i++ {
		randNum := rand.Intn(100)
		originalList[i] = randNum
	}
	setHeap(originalList)
	t.Log(originalList)
}
