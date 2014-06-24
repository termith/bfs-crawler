package queue

import (
	"container/list"
)

//Simple queue based on double-linked list

type Url struct {
	Url   string
	Depth int
}

type UrlQueue struct {
	queue *list.List
}

func NewQueue() *UrlQueue {
	return &UrlQueue{
		queue: list.New(),
	}
}

func (q *UrlQueue) Push(url Url) {
	q.queue.PushBack(url)
}

func (q *UrlQueue) Pop() interface{} {
	if q.queue.Front() == nil {
		return nil
	} else {
		return q.queue.Remove(q.queue.Front())
	}
}

func (q *UrlQueue) Clear() {
	q.queue.Init()
}
