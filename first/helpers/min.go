package helpers

func Min(slice []int) int {
	minVal := slice[0]
	for _, val := range slice {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func MinKey(m map[string]interface{}, f func(string, []int) int) string {
	var minK string
	var minV int
	first := true
	for k, v := range m {
		val := f(k, v.([]int))
		if first || val < minV {
			first = false
			minV = val
			minK = k
		}
	}
	return minK
}
