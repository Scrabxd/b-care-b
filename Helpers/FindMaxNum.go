package helpers

func findMaxValue(m map[string]int) string {
	var maxKey string
	maxValue := 0

	for key, value := range m {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}

	return maxKey
}
