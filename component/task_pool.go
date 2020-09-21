package component

import (
	"fmt"
)

var worker_num int

type Pool struct {
	worker chan func()
	size   chan bool
}

func NewTaskPool(size int) *Pool {
	return &Pool{
		worker: make(chan func()),
		size:   make(chan bool, size),
	}
}

func (pool *Pool) workerStart(worker_num int, task func()) {
	defer func() { <-pool.size }()
	fmt.Printf("worker number start:%d \n", worker_num)
	for {
		task()
		task = <-pool.worker
	}
}

func (pool *Pool) RunTask(task func()) {
	select {
	case pool.worker <- task:
	case pool.size <- true:
		go pool.workerStart(worker_num, task)
		worker_num++
	}
}
