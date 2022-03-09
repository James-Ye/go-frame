package number

func MIN(vals ...int) int {
	var min int
	for _, val := range vals {
		if min == 0 || val <= min {
			min = val
		}
	}
	return min
}

func MAX(vals ...int) int {
	var max int
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
