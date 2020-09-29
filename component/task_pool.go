package component

import (
	"fmt"
)

type Pool struct {
	worker    chan func()
	size      chan bool
	workerNum int
}

func NewTaskPool(size int) *Pool {
	return &Pool{
		worker: make(chan func()),
		size:   make(chan bool, size),
	}
}

func (pool *Pool) workerStart(workerNum int, task func()) {
	defer func() { <-pool.size }()
	fmt.Printf("worker number start:%d \n", workerNum)
	for {
		task()
		task = <-pool.worker
	}
}

func (pool *Pool) RunTask(task func()) {
	select {
	case pool.worker <- task:
	case pool.size <- true:
		go pool.workerStart(pool.workerNum, task)
		pool.workerNum++
	}
}
