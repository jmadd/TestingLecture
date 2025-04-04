package function_by_function_testing_example

import (
	"sort"
)

func ProcessData(data []int) int {
	return processDataWithDependencyInjection(data, processNumber, calculateSum, calculateAverage, calculateMedian, optimalRepresentation)
}

func processDataWithDependencyInjection(data []int, process func(int) int, sum func([]int) int, average func(int, int) float64, median func([]int) float64, optimal func(float64, float64, func(float64) int) int) int {
	processedData := make([]int, len(data))
	for i, value := range data {
		processedData[i] = process(value)
	}
	sumResult := sum(processedData)
	averageResult := average(sumResult, len(processedData))
	medianResult := median(processedData)
	return optimal(averageResult, medianResult, roundOff)
}

func processNumber(number int) int {
	if number < 0 {
		return number + 3
	}
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
	// this is a critical edge case that testing helped to identify
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return float64(data[n/2-1]+data[n/2]) / 2
	}
	return float64(data[n/2])
}

// this function has logical branching without an if statement
// not caught with line coverage
func optimalRepresentation(average float64, median float64, roundOff func(float64) int) int {
	return max(roundOff(average), int(median), 0)
}

func roundOff(number float64) int {
	return int(number + 0.5)
}
