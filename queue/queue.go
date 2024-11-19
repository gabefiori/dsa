package queue

import "errors"

var ErrEmptyQueue = errors.New("Empty queue")

type Node struct {
	Value any
	Next  *Node
}

type Queue struct {
	front  *Node
	rear   *Node
	length int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(val any) {
	newNode := &Node{Value: val}

	if q.rear != nil {
		q.rear.Next = newNode
	}

	q.rear = newNode

	if q.front == nil {
		q.front = newNode
	}

	q.length++
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQueue
	}

	elem := q.front.Value

	q.front = q.front.Next

	if q.front == nil {
		q.rear = nil
	}

	q.length--

	return elem, nil
}

func (q *Queue) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue) Len() int {
	return q.length
}
