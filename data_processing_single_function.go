package TestingExperimentation

import "sort"

func ProcessDataSingleFunc(data []int) int {
	processedData := make([]int, len(data))
	for i, value := range data {
		if value < 0 {
			processedData[i] = value + 3
		} else {
			processedData[i] = value * 2
		}
	}
	sum := 0
	for _, value := range data {
		sum += value
	}
	average := float64(sum) / float64(len(processedData))

	median := 0.0
	sort.Ints(data)
	n := len(data)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		median = float64(data[n/2-1]+data[n/2]) / 2
	} else {
		median = float64(data[n/2])
	}

	return max(int(average+0.5), int(median), 0)
}
