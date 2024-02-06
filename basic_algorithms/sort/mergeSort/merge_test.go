package mergesort_test

import (
	"math/rand"
	"testing"
)

func setMerge(arr []int) {
	arrLength := len(arr)
	if arrLength > 1 {

		mid := arrLength / 2

		leftArr := make([]int, len(arr[:mid]))
		rightArr := make([]int, len(arr[mid:]))
		copy(leftArr, arr[:mid])
		copy(rightArr, arr[mid:])

		setMerge(leftArr)
		setMerge(rightArr)

		i, j, k := 0, 0, 0

		for i < len(leftArr) && j < len(rightArr) {
			if leftArr[i] < rightArr[j] {
				arr[k] = leftArr[i]
				i++
			} else {
				arr[k] = rightArr[j]
				j++
			}
			k++
		}

		for i < len(leftArr) {
			arr[k] = leftArr[i]
			i++
			k++
		}

		for j < len(rightArr) {
			arr[k] = rightArr[j]
			j++
			k++
		}
	}
}

func TestMerge(t *testing.T) {
	listLength := 27
	originalList := make([]int, listLength)
	for i := 0; i < len(originalList); i++ {
		randNum := rand.Intn(100)
		originalList[i] = randNum
	}
	t.Log(originalList)
	setMerge(originalList)
	t.Log(originalList)
}
