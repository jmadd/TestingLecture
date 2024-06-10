package TestingExperimentation

import (
	"testing"
)

func TestProcessDataNoEdgeCase(t *testing.T) {
	t.Run("WithPositiveNumbers", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		  ProcessData(data)
	})

	t.Run("WithNegativeNumbers", func(t *testing.T) {
		data := []int{-1, -2, -3, -4, -5}
		   ProcessData(data)
	})

	t.Run("WithMixedNumbers", func(t *testing.T) {
		data := []int{-1, 2, -3, 4, -5}
		  ProcessData(data)
	})

	t.Run("WithAnEvenNumberOfElements", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		  ProcessData(data)
	})

}
