package main


func Slice(a []string, nbrs ...int) []string {
	length := len(a)

	// helper to convert negative index to positive
	toIndex := func(n int) int {
		if n < 0 {
			return length + n
		}
		return n
	}

	// only one int → slice from that position to end
	if len(nbrs) == 1 {
		start := toIndex(nbrs[0])
		if start < 0 {
			start = 0
		}
		if start > length {
			return nil
		}
		return a[start:]
	}

	// two ints → slice from start to end
	start := toIndex(nbrs[0])
	end := toIndex(nbrs[1])

	// invalid range → return nil
	if start < 0 {
		start = 0
	}
	if end > length {
		end = length
	}
	if start >= end {
		return nil
	}

	return a[start:end]
}