package workerpool

import (
	"fmt"
	"sync"
)

type OddEven interface {
	Print()
}

type oddEven struct {
	pool Pool
}

func NewOddEven(worker int) OddEven {
	return &oddEven{
		pool: NewPool(worker),
	}
}

func (o *oddEven) Print() {
	o.pool.Start()
	defer o.pool.Wait()
	wg := sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			o.pool.Submit(func() {
				if i%2 == 0 {
					fmt.Println("Even", i)
				} else {
					fmt.Println("Odd", i)
				}
				wg.Done()
			})
		}(i)
	}
	wg.Wait()
}
