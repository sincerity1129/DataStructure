package quicksort_test

import (
	quicksort "algorithm/sort/quickSort"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < 100; i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, quicksort.IsSorted(arr))
	quicksort.QuickSort(arr)
	assert.True(t, quicksort.IsSorted(arr), arr)
	t.Log(arr)
}
