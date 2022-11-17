package sync

import (
	"sync"
	"time"
)

type node struct {
	next *node
	prev *node
	val  string
}

type mapValue struct {
	validBefore time.Time
	qnode       *node
}

type syncLRUSet struct {
	m   sync.Mutex
	set map[string]*mapValue

	secondsToAdd int64
	limit        int

	startQueue *node
	endQueue   *node
}

type SetI interface {
	Has(k string) bool
	Set(k string)
}

func NewSyncLRUSet(limit int, secondsToAdd int64) SetI {
	return &syncLRUSet{
		set:          make(map[string]*mapValue, limit),
		secondsToAdd: secondsToAdd,
		limit:        limit,
	}
}

func (s *syncLRUSet) Has(k string) bool {
	s.m.Lock()
	defer s.m.Unlock()
	return s.hasNoLock(k)
}

func (s *syncLRUSet) hasNoLock(k string) bool {
	if _, ok := s.set[k]; !ok {
		return false
	}
	now := time.Now()
	val := s.set[k]
	if now.After(val.validBefore) {
		delete(s.set, k)
		// delete key from queue
		s.deleteNode(val.qnode)
		return false
	}

	s.insertNodeIntoHead(val.qnode, true)
	return true
}

func (s *syncLRUSet) Set(k string) {
	s.m.Lock()
	defer s.m.Unlock()

	if s.hasNoLock(k) {
		s.set[k].validBefore = time.Now().Add(time.Duration(s.secondsToAdd) * time.Second)
		return
	}

	val := &mapValue{}
	if len(s.set) >= s.limit {
		// if limit is reach delete recently used
		delete(s.set, s.endQueue.val)
		s.popNode(s.endQueue)
	}

	// insert into head key from queue
	qnode := &node{val: k}
	s.insertNodeIntoHead(qnode, false)

	val.qnode = qnode
	val.validBefore = time.Now().Add(time.Duration(s.secondsToAdd) * time.Second)
	s.set[k] = val
}

func (s *syncLRUSet) deleteNode(qnode *node) {
	if qnode == nil {
		return
	}

	if qnode.prev != nil {
		qnode.prev.next = qnode.next
	} else {
		s.startQueue = qnode.next
	}

	if qnode.next != nil {
		qnode.next.prev = qnode.prev
	} else {
		s.endQueue = qnode.prev
	}
}

func (s *syncLRUSet) insertNodeIntoHead(qnode *node, inQueue bool) {
	if s.startQueue == nil {
		s.startQueue = qnode
		s.endQueue = qnode
		return
	}

	if s.startQueue == qnode {
		return
	}

	prevEnd := s.endQueue

	if inQueue {
		s.deleteNode(qnode)
	}

	if prevEnd == qnode {
		s.endQueue = qnode.prev
	}

	prevStart := s.startQueue
	s.startQueue = qnode

	qnode.next = prevStart
	qnode.prev = nil

	prevStart.prev = qnode

}

func (s *syncLRUSet) popNode(qnode *node) {
	if s.startQueue == s.endQueue {
		s.startQueue = nil
		s.endQueue = nil
		return
	}

	nextPrev := qnode.prev
	s.deleteNode(qnode)
	s.endQueue = nextPrev
}
