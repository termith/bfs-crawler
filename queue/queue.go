package queue

import (
	"container/list"
)

//Simple queue based on double-linked list

type UrlQueue struct {
	queue *list.List
}

func NewQueue() *UrlQueue {
	return &UrlQueue{
		queue: list.New(),
	}
}

func (q *UrlQueue) Push(url string) {
	q.queue.PushBack(url)
}

func (q *UrlQueue) Pop() interface{} {
	return q.queue.Remove(q.queue.Front())
}
