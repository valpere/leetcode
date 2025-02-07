package queue

import "errors"

type QueueInterface[T any] interface {
	Len() int
	IsEmpty() bool
	IsFull() bool
	Enqueue(elem T) error
	Peek() (elem T, err error)
	Dequeue() (elem T, err error)
}

type Queue[T any] struct {
	Elements []T
	out      int
	inp      int
	cnt      int
}

func NewQueue[T any](size int) *Queue[T] {
	q := new(Queue[T])
	if size == 0 {
		size = 8
	}
	q.Elements = make([]T, size)
	return q
}

func (q *Queue[T]) Len() int {
	return len(q.Elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return (q.cnt == 0)
}

func (q *Queue[T]) IsFull() bool {
	return (q.cnt == q.Len())
}

func (q *Queue[T]) Enqueue(elem T) error {
	if q.IsFull() {
		return errors.New("is full")
	}

	q.Elements[q.inp] = elem

	q.cnt += 1
	q.inp += 1
	if q.inp >= len(q.Elements) {
		q.inp = 0
	}

	return nil
}

func (q *Queue[T]) Peek() (elem T, err error) {
	if q.IsEmpty() {
		return elem, errors.New("is empty")
	}

	return q.Elements[q.out], nil
}

func (q *Queue[T]) Dequeue() (elem T, err error) {
	if q.IsEmpty() {
		return elem, errors.New("is empty")
	}

	elem = q.Elements[q.out]
	q.cnt -= 1
	q.out += 1
	if q.out >= len(q.Elements) {
		q.out = 0
	}

	return elem, nil
}
