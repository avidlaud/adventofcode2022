package utils

import "errors"

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

func MapSliceNoErr[I any, O any](input []I, mapFunc func(I) O) []O {
	newSlice := make([]O, len(input))
	for i, value := range input {
		newSlice[i] = mapFunc(value)
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

type Stack[T comparable] struct {
	Vals []T
}

func (s *Stack[T]) Push(item T) {
	s.Vals = append(s.Vals, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var nullItem T
		return nullItem, false
	}
	top := s.Vals[len(s.Vals)-1]
	s.Vals = s.Vals[:len(s.Vals)-1]
	return top, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Vals) == 0
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var nullItem T
		return nullItem, false
	}
	return s.Vals[len(s.Vals)-1], true
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x == 0 {
		return 0
	}
	return 1
}

type Queue[T any] struct {
	channel chan T
}

func NewQueue[T any](capacity int) *Queue[T] {
	return &Queue[T]{
		channel: make(chan T, capacity),
	}
}

func (q *Queue[T]) Enqueue(val T) error {
	select {
	case q.channel <- val:
		return nil
	default:
		return errors.New("Queue full")
	}
}

func (q *Queue[T]) Dequeue() (T, bool) {
	select {
	case val := <-q.channel:
		return val, true
	default:
		var nilItem T
		return nilItem, false
	}
}

func (q *Queue[T]) Deq() T {
	select {
	case val := <-q.channel:
		return val
	default:
		var nilItem T
		return nilItem
	}
}
