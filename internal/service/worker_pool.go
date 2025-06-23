package service

import (
	"context"
)

type Task func(ctx context.Context) (interface{}, error)

type taskRequest struct {
	ctx      context.Context
	task     Task
	resultCh chan taskResult
}

type taskResult struct {
	resp interface{}
	err  error
}

type WorkerPool struct {
	taskQueue chan taskRequest
}

// NewWorkerPool now accepts a bufferSize parameter for the taskQueue channel.
func NewWorkerPool(numWorkers int, bufferSize int) *WorkerPool {
	wp := &WorkerPool{
		taskQueue: make(chan taskRequest, bufferSize),
	}
	for i := 0; i < numWorkers; i++ {
		go wp.worker()
	}
	return wp
}

func (wp *WorkerPool) worker() {
	for req := range wp.taskQueue {
		resp, err := req.task(req.ctx)
		select {
		case req.resultCh <- taskResult{resp: resp, err: err}:
		case <-req.ctx.Done():
			// Avoid blocking if the context is already done
		}
	}
}

func (wp *WorkerPool) Submit(ctx context.Context, task Task) (interface{}, error) {
	resultCh := make(chan taskResult, 1)
	wp.taskQueue <- taskRequest{
		ctx:      ctx,
		task:     task,
		resultCh: resultCh,
	}
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-resultCh:
		return res.resp, res.err
	}
}

func (wp *WorkerPool) Close() {
	close(wp.taskQueue)
}
