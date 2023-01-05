package DataSturcture

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0)}
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Pop() T {
	var result T
	if len(s.data) == 0 {
		return result
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack[T]) Peek() T {
	var result T
	if len(s.data) == 0 {
		return result
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = make([]T, 0)
}

func (s *Stack[T]) ToSlice() []T {
	return s.data
}

func (s *Stack[T]) FromSlice(data []T) {
	s.data = data
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *Stack[T]) Print() {
	fmt.Println(s)
}

func (s *Stack[T]) ForEach(f func(item T)) {
	for _, item := range s.data {
		f(item)
	}
}

func (s *Stack[T]) Map(f func(item T) T) *Stack[T] {
	newStack := NewStack[T]()
	for _, item := range s.data {
		newStack.Push(f(item))
	}
	return newStack
}

func (s *Stack[T]) Filter(f func(item T) bool) *Stack[T] {
	newStack := NewStack[T]()
	for _, item := range s.data {
		if f(item) {
			newStack.Push(item)
		}
	}
	return newStack
}

func (s *Stack[T]) Reduce(f func(item1 T, item2 T) T) T {
	var result T
	if len(s.data) == 0 {
		return result
	}
	result = s.data[0]
	for _, item := range s.data[1:] {
		result = f(result, item)
	}
	return result
}

func (s *Stack[T]) Find(f func(item T) bool) T {
	var result T
	for _, item := range s.data {
		if f(item) {
			return item
		}
	}
	return result
}

func (s *Stack[T]) FindIndex(f func(item T) bool) int {
	for i, item := range s.data {
		if f(item) {
			return i
		}
	}
	return -1
}
