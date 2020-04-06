package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	unsorted := []int{99, 55, 33, 67, 9, 5, 431, 999, 8108, 108}

	fmt.Println("Before : ", unsorted)

	HeapSort(unsorted)

	target := []int{5, 9, 33, 55, 67, 99, 108, 431, 999, 8108}

	fmt.Println("After : ", unsorted)

	if !reflect.DeepEqual(unsorted, target) {
		t.Error("HeapSort Result error: ", unsorted)
	}
}
