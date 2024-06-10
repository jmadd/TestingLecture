package TestingExperimentation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessDataFinished(t *testing.T) {
	t.Run("WithPositiveNumbers", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		expected := 6
		result := ProcessData(data)
		assert.Equal(t, result, expected)
	})

	t.Run("WithNegativeNumbers", func(t *testing.T) {
		data := []int{-1, -2, -3, -4, -5}
		expected := 0
		result := ProcessData(data)
		assert.Equal(t, result, expected)
	})

	t.Run("WithMixedNumbers", func(t *testing.T) {
		data := []int{-1, 2, -3, 4, -5}
		expected := 2
		result := ProcessData(data)
		assert.Equal(t, result, expected)
	})

	t.Run("WithAnEvenNumberOfElements", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		expected := 5
		result := ProcessData(data)
		assert.Equal(t, result, expected)
	})

	t.Run("WithEmptySlice", func(t *testing.T) {
		var data []int
		expected := 0
		result := ProcessData(data)
		assert.Equal(t, result, expected)
	})

	t.Run("Returns 0 when average and median are negative", func(t *testing.T) {
		data := []int{-11, -12, -13, -14}
		expected := 0
		result := ProcessData(data)
		assert.Equal(t, expected, result)
	})

	t.Run("Returns average when average and median are positive and average is higher", func(t *testing.T) {
		data := []int{11, 12, 13, 20}
		expected := 28
		result := ProcessData(data)
		assert.Equal(t, expected, result)
	})

	t.Run("Returns average when average and median are positive and median is higher", func(t *testing.T) {
		data := []int{1, 12, 13, 14}
		expected := 25
		result := ProcessData(data)
		assert.Equal(t, expected, result)
	})

}
