package shellsort_test

import (
	"math/rand"
	"testing"
)

func setShell(arr []int, gap int) {
	arrLength := len(arr)
	divide := gap

	for gap > 0 {
		for i := gap; i < arrLength; i++ {
			tmp := arr[i]
			j := i
			for j >= gap && arr[j-gap] > tmp {
				arr[j] = arr[j-gap]
				j -= gap
			}
			arr[j] = tmp
		}
		gap /= divide
	}
}

func TestShell(t *testing.T) {
	listLength := 75
	originalList := make([]int, listLength)
	for i := 0; i < len(originalList); i++ {
		randNum := rand.Intn(100)
		originalList[i] = randNum
	}

	setShell(originalList, 3)

	t.Log(originalList)
}
