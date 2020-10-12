package component

import (
	"proxy-collect/component/logger"
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
	logger.Info("worker number start:", workerNum)
	for {
		task()
		task = <-pool.worker
		if task == nil {
			logger.Success("worker exit:", workerNum)
			break
		}
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

func (pool *Pool) Close() {
	l := len(pool.size)
	for i := 0; i < l; i++ {
		pool.worker <- nil
	}
}
