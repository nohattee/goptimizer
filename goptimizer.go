package goptimizer

func MergeSlice[T any](data ...[]T) []T {
	n := 0
	for i := range data {
		if k := n + len(data[i]); k < n {
			panic("slice length overflows")
		} else {
			n = k
		}
	}

	output := make([]T, 0, n)
	for i := range data {
		output = append(output, data[i]...)
	}

	return output
}
