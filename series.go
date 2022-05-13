package stdlib

// TODO: is this a table valued function? How to return a table valued function?
func generate_series(start int, stop int, rest... []int) []int {
	step := 1
	if len(rest) > 0 {
		step = rest[0]
	}

	var res []int
	for i := start; i < stop; i += step {
		res = append(i)
	}

	return res
}
