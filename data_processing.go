package TestingExperimentation

import (
	"sort"
)

func ProcessData(data []int) int {
	processedData := make([]int, len(data))
	for i, value := range data {
		processedData[i] = processNumber(value)
	}
	sum := calculateSum(processedData)
	average := calculateAverage(sum, len(processedData))
	median := calculateMedian(processedData)
	return optimalRepresentation(average, median)
}

func processNumber(number int) int {
	// trivial example where line coverage is not enough
	if number < 0 { return number + 3 }
	return number * 2
}

func calculateSum(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

func calculateAverage(sum int, count int) float64 {
	return float64(sum) / float64(count)
}

func calculateMedian(data []int) float64 {
	sort.Ints(data)
	n := len(data)
	if n == 0 {
		return 0
	}
	if n%2 == 0 { return float64(data[n/2-1]+data[n/2]) / 2 }
	return float64(data[n/2])
}

// Pulled from "Jira product requirements":
// Should return "roundOff-ed" avg, rounded median divided by 2, or 0
// whichever is biggest.
func optimalRepresentation(average float64, median float64) int {
	//return max(roundOff(average), int(median), 0)
	newAvg := roundOff(average)
	newMedian := int(median)
	if newAvg < 0 && newMedian < 0 {
		return 0
	} else if newAvg < newMedian {
		return newMedian / 2
	} else {
		return newAvg
	}
}

func roundOff(number float64) int {
	return int(number + 0.5)
}
