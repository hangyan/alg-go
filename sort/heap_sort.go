package sort

// from http://en.wikipedia.org/wiki/Heapsort
// iParent = floor((i-1) / 2)
func buildMaxHeap(input []int) {
	for i := (len(input) - 1) / 2; i >= 0; i-- {
		maxHeapify(input, i)
	}
}

func maxHeapify(input []int, position int) {
	size := len(input)
	maximum := position

	leftChild := 2*position + 1
	rightChild := 2*position + 1

	if leftChild < size && input[leftChild] > input[position] {
		maximum = leftChild
	}

	if rightChild < size && input[rightChild] > input[position] {
		maximum = rightChild
	}

	if position != maximum {
		input[position], input[maximum] = input[maximum], input[position]
		// TODO: ?
		maxHeapify(input, maximum)
	}
}

func HeapSort(input []int) {
	buildMaxHeap(input)

	for i := len(input) - 1; i >= 1; i-- {
		input[i], input[0] = input[0], input[i]
		maxHeapify(input[:i-1], 0)
	}
}
