package workerpool

import "sync"

type Pool interface {
	Start()
	Submit(func())
	Wait()
}

type pool struct {
	workers int
	wg      sync.WaitGroup
	tasks   chan func()
}

func NewPool(workers int) Pool {
	return &pool{
		workers: workers,
		wg:      sync.WaitGroup{},
		tasks:   make(chan func()),
	}
}

func (p *pool) Start() {
	for i := 0; i < p.workers; i++ {
		go func() {
			for task := range p.tasks {
				task()
				p.wg.Done()
			}
		}()
	}
}

func (p *pool) Submit(task func()) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *pool) Wait() {
	close(p.tasks)
	p.wg.Wait()
}
