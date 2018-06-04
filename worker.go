package workerpool

import (
	"sync"
)

// Task encapsulates a work item that should go in a work pool.
type Task struct {
	// Err holds an error that occurred during a task. Its result is only
	// meaningful after Run has been called for the pool that holds it.
	Payload interface{}
	F       Work
}
type Work func(interface{})interface{}



// Run runs a Task and does appropriate accounting via a given sync.WorkGroup.
func (t *Task) Run(p *Pool) {
	var ret=t.F(t.Payload)
	select {
	case p.RetChan <-ret:
	}
}

// Pool is a worker group that runs a number of tasks at a configured
// concurrency.
type Pool struct {
	concurrency int
	tasksChan   chan *Task
	RetChan     chan interface{}
	wg          sync.WaitGroup
}

// NewPool initializes a new pool with the given tasks and at the given
// concurrency.
func NewPool(concurrency int) *Pool {
	return &Pool{
		concurrency: concurrency,
		tasksChan:   make(chan *Task,concurrency*2),
		RetChan:     make(chan interface{},concurrency*2),
	}
}

func (p *Pool) Start(){
	for i := 0; i < p.concurrency; i++ {
		p.wg.Add(1)
		go p.work()
	}
}

func (p *Pool) NewTask(f Work,payload interface{}){
	select {
		case p.tasksChan<-&Task{F:f, Payload:payload}:
	}
}

// Run runs all work within the pool and blocks until it's finished.
func (p *Pool) Close() {
	// all workers return
	close(p.tasksChan)
	p.wg.Wait()
	close(p.RetChan)
}

// The work loop for any single goroutine.
func (p *Pool) work() {
	defer p.wg.Done()
	var todo *Task
	for todo= range p.tasksChan {
		todo.Run(p)
	}
}

