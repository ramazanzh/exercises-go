package problem1

import (
	"fmt"
)

type Queue struct {
	items []any
}

func (q *Queue) Enqueue(v any) {
	if q.IsEmpty() {
		q.items = make([]any, 0)
	}
	q.items = append(q.items, v)
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	res := q.items[0]

	q.items = q.items[1:]

	return res, nil
}

func (q *Queue) Peek() (any, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("queue is empty")
	}

	return q.items[0], nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}
