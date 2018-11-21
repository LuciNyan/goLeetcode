package main

import "fmt"

/**
	循环队列
 */

type MyCircularQueue struct {

	head int
	tail int
	elem []int
	size int
	max int
}


/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {

	return MyCircularQueue {
		head: 0,
		tail: 0,
		elem: make([]int, k),
		size: 0,
		max: k,
	}

}


/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) EnQueue(value int) bool {
	flag := q.IsFull()
	if flag {

		return false
	}

	q.tail++
	if q.tail == q.max {

		q.tail = 0
	}
	q.elem[q.tail] = value
	q.size++

	if q.size == 1 {
		q.head = q.tail
	}


	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) DeQueue() bool {
	flag := q.IsEmpty()
	if flag {

		return false
	}

	q.size--
	q.head++
	if q.head == q.max {

		q.head = 0
	}

	return true
}


/** Get the front item from the queue. */
func (q *MyCircularQueue) Front() int {

	flag := q.IsEmpty()
	if flag {

		return -1
	}
	return q.elem[q.head]
}


/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {

	flag := q.IsEmpty()
	if flag {

		return -1
	}

	return q.elem[q.tail]
}


/** Checks whether the circular queue is empty or not. */
func (q *MyCircularQueue) IsEmpty() bool {

	if q.size == 0 {

		return true
	}

	return false
}


/** Checks whether the circular queue is full or not. */
func (q *MyCircularQueue) IsFull() bool {

	if q.size == q.max {

		return true
	}

	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main()  {

	haru := Constructor(5)
	fmt.Println(haru)

}
