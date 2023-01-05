package DataSturcture

import "fmt"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}

func (q *Queue[T]) Push(item T) {
	q.data = append(q.data, item)
}

func (q *Queue[T]) Pop() T {
	var result T
	if len(q.data) == 0 {
		return result
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item
}

func (q *Queue[T]) Peek() T {
	var result T
	if len(q.data) == 0 {
		return result
	}
	return q.data[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}

func (q *Queue[T]) ToSlice() []T {
	return q.data
}

func (q *Queue[T]) FromSlice(data []T) {
	q.data = data
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", q.data)
}

func (q *Queue[T]) Print() {
	fmt.Println(q)
}

func (q *Queue[T]) Reverse() {
	for i, j := 0, len(q.data)-1; i < j; i, j = i+1, j-1 {
		q.data[i], q.data[j] = q.data[j], q.data[i]
	}
}

func (q *Queue[T]) ForEach(f func(T)) {
	for _, item := range q.data {
		f(item)
	}
}

func (q *Queue[T]) Map(f func(T) T) {
	for i, item := range q.data {
		q.data[i] = f(item)
	}
}

func (q *Queue[T]) Filter(f func(T) bool) {
	var data []T
	for _, item := range q.data {
		if f(item) {
			data = append(data, item)
		}
	}
	q.data = data
}

func (q *Queue[T]) Reduce(f func(T, T) T) T {
	var result T
	if len(q.data) == 0 {
		return result
	}
	result = q.data[0]
	for _, item := range q.data[1:] {
		result = f(result, item)
	}
	return result
}

func (q *Queue[T]) Find(f func(T) bool) T {
	var result T
	for _, item := range q.data {
		if f(item) {
			return item
		}
	}
	return result
}

func (q *Queue[T]) FindIndex(f func(T) bool) int {
	for i, item := range q.data {
		if f(item) {
			return i
		}
	}
	return -1
}
