package utils

func MapSlice[I any, O any](input []I, mapFunc func(I) (O, error)) []O {
	newSlice := make([]O, len(input))
	for i, value := range input {
		mappedValue, err := mapFunc(value)
		if err != nil {
			panic("error applying map function " + err.Error())
		}
		newSlice[i] = mappedValue
	}
	return newSlice
}

func MaxIntSlice(values []int) (int, int) {
	if len(values) == 0 {
		panic("no elements when attempting to find max")
	}
	maxValue := values[0]
	maxIdx := 0
	for i, value := range values {
		if value > maxValue {
			maxValue = value
			maxIdx = i
		}
	}
	return maxIdx, maxValue
}

func SumSlice(values []int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}
