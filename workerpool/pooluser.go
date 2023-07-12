package workerpool

import (
	"fmt"
	"sync"
)

var ki int

type PoolUser interface {
	Execute()
}

type poolUser struct {
	pool Pool
	wg   sync.WaitGroup
}

func NewPoolUser() PoolUser {
	return &poolUser{
		pool: NewPool(10),
		wg:   sync.WaitGroup{},
	}
}

func (p *poolUser) Execute() {
	lock := sync.Mutex{}
	p.pool.Start()
	defer p.pool.Wait()
	p.wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(lock *sync.Mutex) {
			defer p.wg.Done()
			lock.Lock()
			ki++
			p.pool.Submit(func() {
				count(lock)
			})
		}(&lock)
	}
	p.wg.Wait()
	fmt.Println("total", ki)
}

func count(lock *sync.Mutex) {
	fmt.Println(ki)
	lock.Unlock()
}
