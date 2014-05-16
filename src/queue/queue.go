// Package of queue container. Implement First-In-First-Out queue.
package queue

import (
    "sync"
    "time"
    "errors"
    "container/list"
)

// queue.Queue Interface
type Queue interface {
    // push an item in the Queue back
    Push(interface{})

    // pop an item from the Queue front
    Pop() interface{}

    // return true if there is no item
    Empty() bool

    // clear all items
    Clear()

    // length of queue
    Len() int
}


const (
    DEFAULT_QUEUE_CAP = 64
    POP_TIME_LIMIT = 10
)

var TimeOutErr = errors.New("Time out")

// A Multi-Routine Queue
type MultiRoutineQueue struct {
    data *list.List
    mutex sync.Mutex
    cond *sync.Cond
}

func NewMultiRoutineQueue() *MultiRoutineQueue {
    q := &MultiRoutineQueue{ data : list.New() }
    q.cond = sync.NewCond(&q.mutex)
    return q
}

func NewMultiRoutineQueueWith(_data []interface{}) *MultiRoutineQueue {
    q := &MultiRoutineQueue{ data : list.New() }
    q.cond = sync.NewCond(&q.mutex)
    for _, v := range _data {
        q.Push(v)
    }
    return q
}

func (q *MultiRoutineQueue) Push(item interface{}) {
    q.mutex.Lock()
    q.data.PushBack(item)
    q.cond.Broadcast()
    q.mutex.Unlock()
}

func (q *MultiRoutineQueue) Pop() interface{} {
    if q.Empty() {
        q.mutex.Lock()
        q.cond.Wait()
        q.mutex.Unlock()
    }
    q.mutex.Lock()
    elem := q.data.Front()
    item := q.data.Remove(elem)
    q.mutex.Unlock()
    return item
}

func (q *MultiRoutineQueue) PopWithTimeOut() (item interface{}, err error) {
    if q.Empty() {
        timeout := make(chan bool, 1)
        go func() {
            time.Sleep(POP_TIME_LIMIT * time.Second)
            timeout <- true
        }()
        for {
            q.mutex.Lock()
            q.cond.Wait()
            q.mutex.Unlock()
            select {
            case <-timeout:
                return nil, TimeOutErr
            }
        }
    }
    q.mutex.Lock()
    elem := q.data.Front()
    item = q.data.Remove(elem)
    q.mutex.Unlock()
    return item, nil
}

func (q *MultiRoutineQueue) Empty() bool {
    q.mutex.Lock()
    isEmpty := (q.data.Len() == 0)
    q.mutex.Unlock()
    return isEmpty
}

func (q *MultiRoutineQueue) Clear() {
    q.mutex.Lock()
    q.data = q.data.Init()
    q.mutex.Unlock()
}

func (q *MultiRoutineQueue) Len() int {
    q.mutex.Lock()
    l := q.data.Len()
    q.mutex.Unlock()
    return l
}

