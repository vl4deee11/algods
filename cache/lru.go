package cache

import (
	"fmt"
	"time"
)

type node struct {
	next *node
	prev *node
	key int
	val int
}

func printQueue(queue *node) {
	if queue == nil {
		fmt.Println("queue is empty")
		return
	}
	ptr := queue
	fmt.Printf("node %d ", ptr.key)
	for ptr.next != nil {
		time.Sleep(500000000)
		if ptr.next.prev != ptr {
			fmt.Println("AAAAAAA")
		}
		ptr = ptr.next
		fmt.Printf("node %d ", ptr.key)
	}
	fmt.Println()
}

func LruCached(f func(int) int, cacheSize int) func(int) int {
	// prepare data structures
	var queue *node = nil
	queueLength := 0
	cacheMap := make(map[int]*node)

	decorated := func(arg int) int {
		if val, ok := cacheMap[arg] ; ok {
			// updating queue to set arg as most recently used
			// if it's first element we don't need to extract and pushFront
			if val.prev == nil {
				return val.val
			}
			// extract node from queue
			val.prev.next = val.next
			if val.next != nil {
				val.next.prev = val.prev
			}
			// insert node to the front
			val.prev = nil
			queue = pushFront(queue, val)
			return val.val
		}
		//	here goes cache miss
		newNode := &node{
			next: nil,
			prev: nil,
			key: arg,
			val: f(arg),
		}
		cacheMap[arg] = newNode

		// empty queue
		if queue == nil {
			queue = newNode
			queueLength++
			return newNode.val
		}

		// non-filled cache
		if queueLength < cacheSize {
			queue = pushFront(queue, newNode)
			queueLength++
			return newNode.val
		}

		// full cache, need to discard last
		last := getLast(queue)
		delete(cacheMap, last.key)
		// one element in queue
		if last.prev == nil {
			queue = newNode
			return newNode.val
		}
		last.prev.next = nil
		queue = pushFront(queue, newNode)
		return newNode.val
	}
	return decorated
}

func getLast(queue *node) *node {
	last := queue
	// skip to the end of the queue
	for last.next != nil {
		last = last.next
	}
	return last
}

func pushFront(queue *node, newNode *node) *node {
	newNode.next = queue
	queue.prev = newNode
	return newNode
}