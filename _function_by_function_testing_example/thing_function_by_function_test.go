package _function_by_function_testing_example

import (
	"math"
	"testing"

	"github.com/stretchr/testify/mock"
)

// TODO consider whether there are some interesting examples where a test in this file would catch a bug that the public API tests would not

type MockProcessNumber struct {
	mock.Mock
}

func (m *MockProcessNumber) Execute(number int) int {
	args := m.Called(number)
	return args.Int(0)
}

type MockCalculateSum struct {
	mock.Mock
}

func (m *MockCalculateSum) Execute(data []int) int {
	args := m.Called(data)
	return args.Int(0)
}

type MockCalculateAverage struct {
	mock.Mock
}

func (m *MockCalculateAverage) Execute(sum int, count int) float64 {
	args := m.Called(sum, count)
	return args.Get(0).(float64)
}

type MockCalculateMedian struct {
	mock.Mock
}

func (m *MockCalculateMedian) Execute(data []int) float64 {
	args := m.Called(data)
	return args.Get(0).(float64)
}

type MockOptimalRepresentation struct {
	mock.Mock
}

func (m *MockOptimalRepresentation) Execute(average float64, median float64, roundOff func(float642 float64) int) int {
	args := m.Called(average, median, roundOff)
	return args.Int(0)
}

func TestProcessDataWithDependencyInjection(t *testing.T) {
	t.Run("Dependencies Are Called Appropriately", func(t *testing.T) {
		// both of these lists contain arbitrary numbers
		data := []int{1, 2, 3, 4, 5}
		dataAfterProcessNumber := []int{2, -1, 6, 8, -10}
		expected := 200

		mockProcessNumber := new(MockProcessNumber)
		// TODO what's the method to use where you don't need to pass in "execute"?
		for i, _ := range data {
			mockProcessNumber.On("Execute", data[i]).Return(dataAfterProcessNumber[i])
		}

		mockCalculateSum := new(MockCalculateSum)
		mockCalculateSum.On("Execute", dataAfterProcessNumber).Return(10)

		mockCalculateAverage := new(MockCalculateAverage)
		mockCalculateAverage.On("Execute", 10, len(data)).Return(2.1)

		mockCalculateMedian := new(MockCalculateMedian)
		mockCalculateMedian.On("Execute", dataAfterProcessNumber).Return(3.0)

		mockOptimalRepresentation := new(MockOptimalRepresentation)
		mockOptimalRepresentation.On("Execute", 2.1, 3.0, MockRoundOff{mock.Mock{}}).Return(expected)

		result := processDataWithDependencyInjection(data, mockProcessNumber.Execute, mockCalculateSum.Execute, mockCalculateAverage.Execute, mockCalculateMedian.Execute, mockOptimalRepresentation.Execute)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}

func TestProcessNumber(t *testing.T) {
	t.Run("PositiveNumber", func(t *testing.T) {
		number := 5
		expected := 10
		result := processNumber(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("NegativeNumber", func(t *testing.T) {
		number := -5
		expected := -2
		result := processNumber(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}

func TestCalculateSum(t *testing.T) {
	t.Run("WithPositiveNumbers", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		expected := 15
		result := calculateSum(data)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithNegativeNumbers", func(t *testing.T) {
		data := []int{-1, -2, -3, -4, -5}
		expected := -15
		result := calculateSum(data)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithMixedNumbers", func(t *testing.T) {
		data := []int{-1, 2, -3, 4, -5}
		expected := -3
		result := calculateSum(data)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithEmptySlice", func(t *testing.T) {
		data := []int{}
		expected := 0
		result := calculateSum(data)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}

func TestCalculateAverage(t *testing.T) {
	t.Run("WithPositiveSumAndCount", func(t *testing.T) {
		sum := 15
		count := 5
		expected := 3.0
		result := calculateAverage(sum, count)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithZeroSumAndPositiveCount", func(t *testing.T) {
		sum := 0
		count := 5
		expected := 0.0
		result := calculateAverage(sum, count)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithNegativeSumAndPositiveCount", func(t *testing.T) {
		sum := -15
		count := 5
		expected := -3.0
		result := calculateAverage(sum, count)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithPositiveSumAndZeroCount", func(t *testing.T) {
		sum := 15
		count := 0
		expected := math.Inf(1)
		result := calculateAverage(sum, count)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})
}

func TestCalculateMedian(t *testing.T) {
	t.Run("WithOddNumberOfElements", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5}
		expected := 3.0
		result := calculateMedian(data)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithEvenNumberOfElements", func(t *testing.T) {
		data := []int{1, 2, 3, 4}
		expected := 2.5
		result := calculateMedian(data)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithSingleElement", func(t *testing.T) {
		data := []int{1}
		expected := 1.0
		result := calculateMedian(data)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})

	t.Run("WithEmptySlice", func(t *testing.T) {
		data := []int{}
		expected := 0.0
		result := calculateMedian(data)
		if result != expected {
			t.Errorf("Expected %f, but got %f", expected, result)
		}
	})
}

type MockRoundOff struct {
	mock.Mock
}

func (m *MockRoundOff) Execute(number float64) int {
	args := m.Called(number)
	return args.Int(0)
}

func TestOptimalRepresentation(t *testing.T) {
	t.Run("AverageAndMedianArePositive", func(t *testing.T) {
		average := 5.0
		median := 3.0
		expected := 5

		mockRoundOff := new(MockRoundOff)
		mockRoundOff.On("Execute", mock.Anything).Return(5)

		result := optimalRepresentation(average, median, mockRoundOff.Execute)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	// Similar test cases would be written for "AverageIsNegativeMedianIsPositive", "AverageIsPositiveMedianIsNegative", "AverageAndMedianAreNegative", and "AverageAndMedianAreZero"
}

func TestRoundOff(t *testing.T) {
	t.Run("WithPositiveNumber", func(t *testing.T) {
		number := 1.4
		expected := 1
		result := roundOff(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithPositiveNumberRoundingUp", func(t *testing.T) {
		number := 1.6
		expected := 2
		result := roundOff(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithNegativeNumber", func(t *testing.T) {
		number := -1.4
		expected := 0
		result := roundOff(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithNegativeNumberRoundingDown", func(t *testing.T) {
		number := -1.6
		expected := -1
		result := roundOff(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})

	t.Run("WithZero", func(t *testing.T) {
		number := 0.0
		expected := 0
		result := roundOff(number)
		if result != expected {
			t.Errorf("Expected %d, but got %d", expected, result)
		}
	})
}
