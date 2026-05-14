package listops

// IntList is an abstraction of a list of integers which we can define methods on.
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	// accumulator = snowball. Start with initial, roll through the list.
	accumulator := initial

	for _, val := range s {
		// Update running total with the next value.
		accumulator = fn(accumulator, val)
	}
	return accumulator
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	accumulator := initial
	// Backward crunch: start at L-1, walk down to 0.
	for i := len(s) - 1; i >= 0; i-- {
		// Flip the inputs: current value enters fn before the accumulator.
		accumulator = fn(s[i], accumulator)
	}
	return accumulator
}

func (s IntList) Filter(fn func(int) bool) IntList {
	// Pre-allocate worst-case (all numbers pass). Prevents mid-loop re-allocs.
	L := len(s)
	result := make([]int, L)

	place := 0 // Manual cursor to track successful matches.
	for _, val := range s {
		if fn(val) {
			result[place] = val
			place++
		}
	}
	// Shrink slice to 'place' to cut off trailing zeros.
	return IntList(result[:place])
}

func (s IntList) Length() int {
	total := 0
	for range s {
		total++
	}
	return total
}

func (s IntList) Map(fn func(int) int) IntList {
	// 1-to-1 transformation. Allocation size is guaranteed.
	L := len(s)
	result := make([]int, L)

	for i, val := range s {
		// Mirror the index directly. No place marker needed.
		result[i] = fn(val)
	}
	return IntList(result)
}

func (s IntList) Reverse() IntList {
	L := len(s)
	result := make([]int, L)

	for i, number := range s {
		// Mirror logic: index 0 goes to L-1-0, etc.
		result[L-1-i] = number
	}
	return IntList(result)
}

func (s IntList) Append(lst IntList) IntList {
	// Sum lengths for a single allocation.
	L := len(s) + len(lst)
	result := make([]int, L)

	for i, number := range s {
		result[i] = number
	}

	for i, number := range lst {
		// Offset second list by the length of the first.
		result[i+len(s)] = number
	}
	return IntList(result)
}

func (s IntList) Concat(lists []IntList) IntList {
	// Calculate total footprint across all nested lists first.
	totalLen := len(s)
	for _, list := range lists {
		totalLen += len(list)
	}

	result := make([]int, totalLen)
	place := 0 // Global cursor for the flat result list.

	// Copy base list.
	for _, val := range s {
		result[place] = val
		place++
	}

	// Double loop to unpack and flatten.
	for _, list := range lists {
		for _, val := range list {
			result[place] = val
			place++
		}
	}
	return IntList(result)
}