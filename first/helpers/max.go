package helpers

func Max(slice []int) int {
	maxVal := slice[0]
	for _, val := range slice {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}

func MaxKey(m map[string]interface{}, f func(string, []int) int) string {
	var maxK string
	var maxV int
	first := true
	for k, v := range m {
		val := f(k, v.([]int))
		if first || val > maxV {
			first = false
			maxV = val
			maxK = k
		}
	}
	return maxK
}
